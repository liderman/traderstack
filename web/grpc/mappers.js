import { UpdateRequest } from "~/grpc/gen/js/liderman/traderstack/stack/v1/stack_api_pb";
import { SetItem, SetStackFunc, SetArgument, Variable, Value } from "~/grpc/gen/js/liderman/traderstack/stack/v1/stack_pb";

const stackToUpdateRequest = (stack, stackId, stackName) => {
  const req = new UpdateRequest();
  req.setId(stackId);
  req.setName(stackName);
  stack.forEach((elItem, indexItem) => {
    const item = new SetItem();
    item.setVariable(elItem.variable);
    item.setStackFunc(mapStackFunc(elItem.stackFunc));

    req.addItems(item, indexItem);
  });

  return req;
}

const mapStackFunc = (stackFunc) => {
  if (!stackFunc) {
    return undefined;
  }

  const ret = new SetStackFunc();
  ret.setName(stackFunc.name);
  stackFunc.argumentsList.forEach((argument, indexArg) => {
    ret.addArguments(mapArgument(argument), indexArg);
  })

  return ret;
}

const mapArgument = (argument) => {
  const ret = new SetArgument();
  ret.setId(argument.id);
  if (argument.variable !== undefined) {
    const variable = new Variable();
    variable.setName(argument.variable.name ? argument.variable.name  : '');
    ret.setVariable(variable);
    return ret;
  }

  const value = new Value();
  switch (argument.baseType) {
    case "string":
      value.setString(argument.input.string);
      break;
    case "integer":
      value.setInteger(argument.input.integer);
      break
    case "decimal":
      value.setDecimal(argument.input.decimal);
      break
    case "boolean":
      value.setBoolean(argument.input.pb_boolean);
      break
  }

  ret.setInput(value);

  return ret;
}

const setValueByArgument = (argument, varType, val) => {
  const ret = {
    input: undefined,
    variable: undefined,
  };

  if (varType === 'variable') {
    const variable = new Variable();
    variable.setName(val === undefined ? '' : val);
    ret.variable = variable.toObject();
    return ret;
  }

  const value = new Value();
  switch (argument.baseType) {
    case "string":
      value.setString(val);
      break;
    case "integer":
      value.setInteger(val);
      break;
    case "decimal":
      value.setDecimal(val);
      break;
    case "boolean":
      value.setBoolean(val);
      break;
  }

  ret.input = value.toObject();
  return ret;
}

export {
  stackToUpdateRequest,
  setValueByArgument,
}
