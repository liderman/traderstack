# TraderStack
> Инструмент помогающий писать торговые стратегии используя визуальный редактор.
> Работает на API ТИНЬКОФФ Инвестиций

## Инструкция по запуску
> appName - traderstack
1. Сборка docker image (не обязательно. Образ опубликован на Docker Hub)
    ```shell
    git clone https://github.com/liderman/traderstack.git
    cd traderstack
    docker build -t "liderman/traderstack:latest" .
    ```

2. Запустите приложение
    > (!) ВАЖНО: не забудьте заменить `ТОКЕН_ПЕСОЧНИЦЫ` и `БОЕВОЙ_ТОКЕН` на свои.
    > 
    > Если вы переживаете указывать боевой токен, то можете указать вместо него токен
    > песочницы, НО в таком режиме будут выставляться ордера только в режиме тестирования стека,
    > а в режиме работы стратегии будут ошибки.
    ```shell
    docker run --rm -p9000:8080 \
     --name traderstack \
     -e TS_BROKER_SANDBOX_API_TOKEN=ТОКЕН_ПЕСОЧНИЦЫ \
     -e TS_BROKER_API_TOKEN=БОЕВОЙ_ТОКЕН \
     liderman/traderstack
    ```

3. Откройте TraderStack в браузере: http://127.0.0.1:9000/
4. Создайте свой первый Стек

    `Стеки > Создать стек`
    
    Введите название `Стратегия RSI` и нажмите `Создать...`
5. Используя панель инструментов создайте следующий сценарий
   ```
   $figi              : FigiByTicker(TCS)
   $lowerRsiThreshold : Integer(30)
   $takeProfit        : Decimal(15)
   $stopLoss          : Decimal(5)
   $lots              : Integer(1)
   $rsi               : RSI($figi, 14)
   $signalBuy         : <=($rsi, $lowerRsiThreshold)
   $lotsProfile       : PortfolioLots($figi)
   $buyResult         : ActionBuyMarket($signalBuy, $figi, $lots)
   $takeProfitResult  : ActionTakeProfit($figi, $takeProfit)
   $stopLossResult    : ActionStopLoss($figi, $stopLoss)
   ```
6. Нажмите `Сохранить` и нажмите `Тест`, чтобы протестировать Стек.
   
    В окне `Протестировать стек` выберите `Дату и время` за которую
    подгрузятся данные и счёт (песочницы).

    Если счёта нет, то не беда, зайдите `Песочница > Счета > Создать счёт`.
    Там же можно и пополнить его.
7. Настало время превратить Стек в Стратегию.

    Откройте `Стратегии > Создать стратегию` и заполните поля:
    * Аккаунт (тут уже реальный аккаунт на бирже)
    * Стек - `Стратегия RSI`
    * Периодичность запуска - `86400` (будет запускаться раз в сутки, но можно указать любое другое число)
    * Стратегия - переключите во `Включено`

    Нажмите `Сохранить`
8. Отслеживайте результаты стратегии
    
    В карточке стратегии `Логи > Обновить` - отобразятся логи.
    В логах можно увидеть запуски, ошибки, попытки исполнения действий


## Для разработчиков
Перегенерация gRPC-клиентов на основе proto-фалов:
```shell
docker run --rm -v $PWD/internal/grpc:/data namely/protoc-all \
-i /data/proto/tinkoff/investapi -d /data/proto/tinkoff/investapi -o /data/gen/tinkoff/investapi -l go
```
Перегенерация gRPC-сервера:
```bash
docker run --rm -v "$(pwd)/internal/grpcsrv:/work" uber/prototool:latest prototool generate
cp -R ./internal/grpcsrv/gen/js/ ./web/grpc/gen/js/
```
Запуск линтера .proto фалов:
```bash
docker run --rm -v "$(pwd)/internal/grpcsrv:/work" uber/prototool:latest prototool lint
```

Запуск в режиме разработки:
```bash
cd cmd/traderstack
go build
./traderstack
```
```bash
cd web
yarn
yarn dev
```
```bash
cd ci/dev
docker-compose up
```
Dev URL: http://127.0.0.1:8080/

## LICENSE
Apache License 2.0

## MAINTAINER
Konstantin Osipov