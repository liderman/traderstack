package engine

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/liderman/traderstack/internal/domain"
	"go.uber.org/zap"
	"strings"
	"time"
)

type StrategyEngine struct {
	sm     *StackManager
	sfr    *StackFuncRepository
	logger *zap.Logger
	data   []*StrategyContext
}

func NewStrategyEngine(sm *StackManager, sfr *StackFuncRepository, logger *zap.Logger) *StrategyEngine {
	return &StrategyEngine{
		sm:     sm,
		sfr:    sfr,
		logger: logger,
	}
}

type StrategyContext struct {
	Strategy  *domain.Strategy
	LastRun   time.Time
	Logs      []*domain.StrategyLog
	IsRunning bool
}

func (s *StrategyContext) AllowRun(now time.Time) bool {
	return s.LastRun.Add(s.Strategy.RunInterval).Before(now) && !s.IsRunning && s.Strategy.Enabled
}

func (s *StrategyContext) WriteErrorLog(text string) {
	s.Logs = append(s.Logs, &domain.StrategyLog{
		LogType: domain.StrategyLogTypeError,
		Message: text,
		Time:    time.Now(),
	})
}

func (s *StrategyContext) WriteStartLog() {
	s.Logs = append(s.Logs, &domain.StrategyLog{
		LogType: domain.StrategyLogTypeStart,
		Message: fmt.Sprintf("Запуск стратегии '%s'", s.Strategy.Id),
		Time:    time.Now(),
	})
}

func (s *StrategyEngine) Create() *domain.Strategy {
	strategy := &domain.Strategy{
		Id: uuid.New().String(),
	}
	s.data = append(s.data, &StrategyContext{
		Strategy:  strategy,
		Logs:      []*domain.StrategyLog{},
		IsRunning: false,
	})

	return strategy
}

func (s *StrategyEngine) Get(id string) *domain.Strategy {
	for _, v := range s.data {
		if v.Strategy.Id == id {
			return v.Strategy
		}
	}

	return nil
}

func (s *StrategyEngine) GetLogs(id string) []*domain.StrategyLog {
	for _, v := range s.data {
		if v.Strategy.Id == id {
			return v.Logs
		}
	}

	return nil
}

func (s *StrategyEngine) GetAll() []*domain.Strategy {
	var ret []*domain.Strategy
	for _, v := range s.data {
		ret = append(ret, v.Strategy)
	}
	return ret
}

func (s *StrategyEngine) Update(strategy *domain.Strategy) error {
	var ctx *StrategyContext
	for _, v := range s.data {
		if v.Strategy.Id == strategy.Id {
			ctx = v
			break
		}
	}
	if ctx == nil {
		return errors.New("strategy is not exist")
	}

	ctx.Strategy = strategy
	return nil
}

func (s *StrategyEngine) Delete(id string) {
	for idx, v := range s.data {
		if v.Strategy.Id == id {
			s.data = append(s.data[:idx], s.data[idx+1:]...)
			break
		}
	}
}

func (s *StrategyEngine) Run() {
	for true {
		s.runStrategies(time.Now())
		time.Sleep(time.Second)
	}
}

func (s *StrategyEngine) runStrategies(now time.Time) {
	for _, strategy := range s.data {
		if !strategy.AllowRun(now) {
			continue
		}

		go s.runStrategy(strategy, now)
	}
}

func (s *StrategyEngine) runStrategy(strategy *StrategyContext, now time.Time) {
	strategy.IsRunning = true
	strategy.LastRun = now
	defer func() {
		strategy.IsRunning = false
	}()

	strategy.WriteStartLog()
	s.logger.Info("Run strategy", zap.String("strategy", strategy.Strategy.Id))

	stack := s.sm.Get(strategy.Strategy.StackId)
	if stack == nil {
		s.logger.Warn("Strategy is not found", zap.String("strategy", strategy.Strategy.Id))
		strategy.WriteErrorLog(
			fmt.Sprintf("Стек '%s' не найден", strategy.Strategy.StackId),
		)
		return
	}

	s.runStack(stack, strategy)
}

func (s *StrategyEngine) runStack(stack *Stack, ctx *StrategyContext) {
	varIdx := map[string]interface{}{}
	for _, v := range stack.Items {
		varIdx = s.runStackItem(v, varIdx, ctx)
	}
}

func (s *StrategyEngine) runStackItem(item *StackItem, varIdx map[string]interface{}, ctx *StrategyContext) map[string]interface{} {
	if item.StackFunc == nil {
		ctx.WriteErrorLog(fmt.Sprintf("Не задана функция для элемента стека '%s'", item.Variable))
		return varIdx
	}

	fnc := s.sfr.Get(item.StackFunc.Name)
	if fnc == nil {
		ctx.WriteErrorLog(fmt.Sprintf("Функция '%s' не существует", item.StackFunc.Name))
		return varIdx
	}

	options := NewOptions(item.StackFunc.Arguments, varIdx)
	value, err := fnc.Run(options, time.Now(), ctx.Strategy.AccountId, false)
	if err != nil {
		ctx.WriteErrorLog(fmt.Sprintf("Ошибка выполнения функции '%s': %s", item.StackFunc.Name, err))
		return varIdx
	}

	s.writeSuccessLog(ctx, value, item.StackFunc)
	varIdx[item.Variable] = value

	return varIdx
}

func (s *StrategyEngine) writeSuccessLog(ctx *StrategyContext, value interface{}, fnc *StackFunc) {
	isActionFunc := strings.HasPrefix(strings.ToLower(fnc.Name), "action")
	if !isActionFunc {
		return
	}

	result, ok := value.(bool)
	if !ok || !result {
		return
	}

	ctx.Logs = append(ctx.Logs, &domain.StrategyLog{
		LogType: domain.StrategyLogTypeAction,
		Message: fmt.Sprintf("Сработало событие '%s'", fnc.Name),
		Time:    time.Now(),
	})
}
