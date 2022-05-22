/**
 * @fileoverview gRPC-Web generated client stub for liderman.traderstack.strategy.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var liderman_traderstack_strategy_v1_strategy_pb = require('../../../../liderman/traderstack/strategy/v1/strategy_pb.js')
const proto = {};
proto.liderman = {};
proto.liderman.traderstack = {};
proto.liderman.traderstack.strategy = {};
proto.liderman.traderstack.strategy.v1 = require('./strategy_api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.strategy.v1.CreateRequest,
 *   !proto.liderman.traderstack.strategy.v1.CreateResponse>}
 */
const methodInfo_StrategyAPI_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.strategy.v1.CreateResponse,
  /** @param {!proto.liderman.traderstack.strategy.v1.CreateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.strategy.v1.CreateResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.strategy.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.strategy.v1.CreateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.strategy.v1.CreateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Create',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Create,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.strategy.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.strategy.v1.CreateResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Create',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Create);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.strategy.v1.GetRequest,
 *   !proto.liderman.traderstack.strategy.v1.GetResponse>}
 */
const methodInfo_StrategyAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.strategy.v1.GetResponse,
  /** @param {!proto.liderman.traderstack.strategy.v1.GetRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.strategy.v1.GetResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.strategy.v1.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.strategy.v1.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.strategy.v1.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Get',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Get,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.strategy.v1.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.strategy.v1.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Get',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.strategy.v1.GetAllRequest,
 *   !proto.liderman.traderstack.strategy.v1.GetAllResponse>}
 */
const methodInfo_StrategyAPI_GetAll = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.strategy.v1.GetAllResponse,
  /** @param {!proto.liderman.traderstack.strategy.v1.GetAllRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.strategy.v1.GetAllResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.strategy.v1.GetAllRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.strategy.v1.GetAllResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.strategy.v1.GetAllResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIClient.prototype.getAll =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/GetAll',
      request,
      metadata || {},
      methodInfo_StrategyAPI_GetAll,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.strategy.v1.GetAllRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.strategy.v1.GetAllResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIPromiseClient.prototype.getAll =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/GetAll',
      request,
      metadata || {},
      methodInfo_StrategyAPI_GetAll);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.strategy.v1.UpdateRequest,
 *   !proto.liderman.traderstack.strategy.v1.UpdateResponse>}
 */
const methodInfo_StrategyAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.strategy.v1.UpdateResponse,
  /** @param {!proto.liderman.traderstack.strategy.v1.UpdateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.strategy.v1.UpdateResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.strategy.v1.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.strategy.v1.UpdateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.strategy.v1.UpdateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Update',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Update,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.strategy.v1.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.strategy.v1.UpdateResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Update',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.strategy.v1.DeleteRequest,
 *   !proto.liderman.traderstack.strategy.v1.DeleteResponse>}
 */
const methodInfo_StrategyAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.strategy.v1.DeleteResponse,
  /** @param {!proto.liderman.traderstack.strategy.v1.DeleteRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.strategy.v1.DeleteResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.strategy.v1.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.strategy.v1.DeleteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.strategy.v1.DeleteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Delete',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Delete,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.strategy.v1.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.strategy.v1.DeleteResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/Delete',
      request,
      metadata || {},
      methodInfo_StrategyAPI_Delete);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.strategy.v1.GetLogsRequest,
 *   !proto.liderman.traderstack.strategy.v1.GetLogsResponse>}
 */
const methodInfo_StrategyAPI_GetLogs = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.strategy.v1.GetLogsResponse,
  /** @param {!proto.liderman.traderstack.strategy.v1.GetLogsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.strategy.v1.GetLogsResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.strategy.v1.GetLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.strategy.v1.GetLogsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.strategy.v1.GetLogsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIClient.prototype.getLogs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/GetLogs',
      request,
      metadata || {},
      methodInfo_StrategyAPI_GetLogs,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.strategy.v1.GetLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.strategy.v1.GetLogsResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.strategy.v1.StrategyAPIPromiseClient.prototype.getLogs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.strategy.v1.StrategyAPI/GetLogs',
      request,
      metadata || {},
      methodInfo_StrategyAPI_GetLogs);
};


module.exports = proto.liderman.traderstack.strategy.v1;

