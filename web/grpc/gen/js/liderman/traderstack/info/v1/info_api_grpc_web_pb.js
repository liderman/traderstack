/**
 * @fileoverview gRPC-Web generated client stub for liderman.traderstack.info.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var liderman_traderstack_info_v1_info_pb = require('../../../../liderman/traderstack/info/v1/info_pb.js')
const proto = {};
proto.liderman = {};
proto.liderman.traderstack = {};
proto.liderman.traderstack.info = {};
proto.liderman.traderstack.info.v1 = require('./info_api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.liderman.traderstack.info.v1.InfoAPIClient =
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
proto.liderman.traderstack.info.v1.InfoAPIPromiseClient =
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
 *   !proto.liderman.traderstack.info.v1.SearchInstrumentRequest,
 *   !proto.liderman.traderstack.info.v1.SearchInstrumentResponse>}
 */
const methodInfo_InfoAPI_SearchInstrument = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.info.v1.SearchInstrumentResponse,
  /** @param {!proto.liderman.traderstack.info.v1.SearchInstrumentRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.info.v1.SearchInstrumentResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.info.v1.SearchInstrumentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.info.v1.SearchInstrumentResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.info.v1.SearchInstrumentResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.info.v1.InfoAPIClient.prototype.searchInstrument =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/SearchInstrument',
      request,
      metadata || {},
      methodInfo_InfoAPI_SearchInstrument,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.info.v1.SearchInstrumentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.info.v1.SearchInstrumentResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.info.v1.InfoAPIPromiseClient.prototype.searchInstrument =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/SearchInstrument',
      request,
      metadata || {},
      methodInfo_InfoAPI_SearchInstrument);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.info.v1.AccountsRequest,
 *   !proto.liderman.traderstack.info.v1.AccountsResponse>}
 */
const methodInfo_InfoAPI_Accounts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.info.v1.AccountsResponse,
  /** @param {!proto.liderman.traderstack.info.v1.AccountsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.info.v1.AccountsResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.info.v1.AccountsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.info.v1.AccountsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.info.v1.AccountsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.info.v1.InfoAPIClient.prototype.accounts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/Accounts',
      request,
      metadata || {},
      methodInfo_InfoAPI_Accounts,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.info.v1.AccountsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.info.v1.AccountsResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.info.v1.InfoAPIPromiseClient.prototype.accounts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/Accounts',
      request,
      metadata || {},
      methodInfo_InfoAPI_Accounts);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.info.v1.SandboxAccountsRequest,
 *   !proto.liderman.traderstack.info.v1.SandboxAccountsResponse>}
 */
const methodInfo_InfoAPI_SandboxAccounts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.info.v1.SandboxAccountsResponse,
  /** @param {!proto.liderman.traderstack.info.v1.SandboxAccountsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.info.v1.SandboxAccountsResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.info.v1.SandboxAccountsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.info.v1.SandboxAccountsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.info.v1.SandboxAccountsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.info.v1.InfoAPIClient.prototype.sandboxAccounts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/SandboxAccounts',
      request,
      metadata || {},
      methodInfo_InfoAPI_SandboxAccounts,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.info.v1.SandboxAccountsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.info.v1.SandboxAccountsResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.info.v1.InfoAPIPromiseClient.prototype.sandboxAccounts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/SandboxAccounts',
      request,
      metadata || {},
      methodInfo_InfoAPI_SandboxAccounts);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.info.v1.OpenSandboxAccountRequest,
 *   !proto.liderman.traderstack.info.v1.OpenSandboxAccountResponse>}
 */
const methodInfo_InfoAPI_OpenSandboxAccount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.info.v1.OpenSandboxAccountResponse,
  /** @param {!proto.liderman.traderstack.info.v1.OpenSandboxAccountRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.info.v1.OpenSandboxAccountResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.info.v1.OpenSandboxAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.info.v1.OpenSandboxAccountResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.info.v1.OpenSandboxAccountResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.info.v1.InfoAPIClient.prototype.openSandboxAccount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/OpenSandboxAccount',
      request,
      metadata || {},
      methodInfo_InfoAPI_OpenSandboxAccount,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.info.v1.OpenSandboxAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.info.v1.OpenSandboxAccountResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.info.v1.InfoAPIPromiseClient.prototype.openSandboxAccount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/OpenSandboxAccount',
      request,
      metadata || {},
      methodInfo_InfoAPI_OpenSandboxAccount);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.info.v1.GetSandboxPositionsRequest,
 *   !proto.liderman.traderstack.info.v1.GetSandboxPositionsResponse>}
 */
const methodInfo_InfoAPI_GetSandboxPositions = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.info.v1.GetSandboxPositionsResponse,
  /** @param {!proto.liderman.traderstack.info.v1.GetSandboxPositionsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.info.v1.GetSandboxPositionsResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.info.v1.GetSandboxPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.info.v1.GetSandboxPositionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.info.v1.GetSandboxPositionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.info.v1.InfoAPIClient.prototype.getSandboxPositions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/GetSandboxPositions',
      request,
      metadata || {},
      methodInfo_InfoAPI_GetSandboxPositions,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.info.v1.GetSandboxPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.info.v1.GetSandboxPositionsResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.info.v1.InfoAPIPromiseClient.prototype.getSandboxPositions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/GetSandboxPositions',
      request,
      metadata || {},
      methodInfo_InfoAPI_GetSandboxPositions);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.liderman.traderstack.info.v1.SandboxPayInRequest,
 *   !proto.liderman.traderstack.info.v1.SandboxPayInResponse>}
 */
const methodInfo_InfoAPI_SandboxPayIn = new grpc.web.AbstractClientBase.MethodInfo(
  proto.liderman.traderstack.info.v1.SandboxPayInResponse,
  /** @param {!proto.liderman.traderstack.info.v1.SandboxPayInRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.liderman.traderstack.info.v1.SandboxPayInResponse.deserializeBinary
);


/**
 * @param {!proto.liderman.traderstack.info.v1.SandboxPayInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.liderman.traderstack.info.v1.SandboxPayInResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.liderman.traderstack.info.v1.SandboxPayInResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.liderman.traderstack.info.v1.InfoAPIClient.prototype.sandboxPayIn =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/SandboxPayIn',
      request,
      metadata || {},
      methodInfo_InfoAPI_SandboxPayIn,
      callback);
};


/**
 * @param {!proto.liderman.traderstack.info.v1.SandboxPayInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.liderman.traderstack.info.v1.SandboxPayInResponse>}
 *     A native promise that resolves to the response
 */
proto.liderman.traderstack.info.v1.InfoAPIPromiseClient.prototype.sandboxPayIn =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/liderman.traderstack.info.v1.InfoAPI/SandboxPayIn',
      request,
      metadata || {},
      methodInfo_InfoAPI_SandboxPayIn);
};


module.exports = proto.liderman.traderstack.info.v1;

