package xerr

const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// SERVER_COMMON_ERROR 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQES_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const GRPC_CODE uint32 = 100

//用户模块
