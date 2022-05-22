import { StackAPIPromiseClient } from "./gen/js/liderman/traderstack/stack/v1/stack_api_grpc_web_pb";
import { InfoAPIPromiseClient } from "./gen/js/liderman/traderstack/info/v1/info_api_grpc_web_pb";
import { StrategyAPIPromiseClient } from "./gen/js/liderman/traderstack/strategy/v1/strategy_api_grpc_web_pb";

const clientStack = new StackAPIPromiseClient('/rpc', null, null);
const clientInfo = new InfoAPIPromiseClient('/rpc', null, null);
const clientStrategy = new StrategyAPIPromiseClient('/rpc', null, null);

export {
  clientStack,
  clientInfo,
  clientStrategy,
}
