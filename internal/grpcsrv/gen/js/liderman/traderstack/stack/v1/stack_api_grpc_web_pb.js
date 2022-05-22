/**
 * @fileoverview gRPC-Web generated client stub for liderman.traderstack.stack.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var liderman_traderstack_stack_v1_stack_pb = require('../../../../liderman/traderstack/stack/v1/stack_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.liderman = {};
proto.liderman.traderstack = {};
proto.liderman.traderstack.stack = {};
proto.liderman.traderstack.stack.v1 = require('./stack_api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.liderman.traderstack.stack.v1.StackAPIClient =
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
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient =
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
 *   !proto.liderman.traderstack.stack.v1.CreateRequest,
 *   !proto.liderman.traderstack.stack.v1.CreateResponse>}
 */
const methodInfo_StackAPI_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.CreateResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.CreateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.CreateResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.CreateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.CreateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Create',
      request,
      metadata || {},
      methodInfo_StackAPI_Create,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.CreateResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Create',
      request,
      metadata || {},
      methodInfo_StackAPI_Create);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.stack.v1.UpdateRequest,
 *   !proto.liderman.traderstack.stack.v1.UpdateResponse>}
 */
const methodInfo_StackAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.UpdateResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.UpdateRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.UpdateResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.UpdateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.UpdateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Update',
      request,
      metadata || {},
      methodInfo_StackAPI_Update,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.UpdateResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Update',
      request,
      metadata || {},
      methodInfo_StackAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.stack.v1.DeleteRequest,
 *   !proto.liderman.traderstack.stack.v1.DeleteResponse>}
 */
const methodInfo_StackAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.DeleteResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.DeleteRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.DeleteResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.DeleteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.DeleteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Delete',
      request,
      metadata || {},
      methodInfo_StackAPI_Delete,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.DeleteResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Delete',
      request,
      metadata || {},
      methodInfo_StackAPI_Delete);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.stack.v1.GetRequest,
 *   !proto.liderman.traderstack.stack.v1.GetResponse>}
 */
const methodInfo_StackAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.GetResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.GetRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.GetResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Get',
      request,
      metadata || {},
      methodInfo_StackAPI_Get,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Get',
      request,
      metadata || {},
      methodInfo_StackAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.stack.v1.GetAllRequest,
 *   !proto.liderman.traderstack.stack.v1.GetAllResponse>}
 */
const methodInfo_StackAPI_GetAll = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.GetAllResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.GetAllRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.GetAllResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.GetAllRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.GetAllResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.GetAllResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.getAll =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/GetAll',
      request,
      metadata || {},
      methodInfo_StackAPI_GetAll,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.GetAllRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.GetAllResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.getAll =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/GetAll',
      request,
      metadata || {},
      methodInfo_StackAPI_GetAll);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.stack.v1.TestRequest,
 *   !proto.liderman.traderstack.stack.v1.TestResponse>}
 */
const methodInfo_StackAPI_Test = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.TestResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.TestRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.TestResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.TestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.TestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.TestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.test =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Test',
      request,
      metadata || {},
      methodInfo_StackAPI_Test,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.TestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.TestResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.test =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/Test',
      request,
      metadata || {},
      methodInfo_StackAPI_Test);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.stack.v1.FuncListRequest,
 *   !proto.liderman.traderstack.stack.v1.FuncListResponse>}
 */
const methodInfo_StackAPI_FuncList = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.FuncListResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.FuncListRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.FuncListResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.FuncListRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.FuncListResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.FuncListResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.funcList =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/FuncList',
      request,
      metadata || {},
      methodInfo_StackAPI_FuncList,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.FuncListRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.FuncListResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.funcList =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/FuncList',
      request,
      metadata || {},
      methodInfo_StackAPI_FuncList);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.stack.v1.FuncArgumentVarListRequest,
 *   !proto.liderman.traderstack.stack.v1.FuncArgumentVarListResponse>}
 */
const methodInfo_StackAPI_FuncArgumentVarList = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.stack.v1.FuncArgumentVarListResponse,
  /** @param {!proto.liderman.traderstack.stack.v1.FuncArgumentVarListRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.stack.v1.FuncArgumentVarListResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.stack.v1.FuncArgumentVarListRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.stack.v1.FuncArgumentVarListResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.stack.v1.FuncArgumentVarListResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.stack.v1.StackAPIClient.prototype.funcArgumentVarList =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/FuncArgumentVarList',
      request,
      metadata || {},
      methodInfo_StackAPI_FuncArgumentVarList,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.stack.v1.FuncArgumentVarListRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.stack.v1.FuncArgumentVarListResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.stack.v1.StackAPIPromiseClient.prototype.funcArgumentVarList =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.stack.v1.StackAPI/FuncArgumentVarList',
      request,
      metadata || {},
      methodInfo_StackAPI_FuncArgumentVarList);
};


module.exports = proto.liderman.traderstack.stack.v1;

