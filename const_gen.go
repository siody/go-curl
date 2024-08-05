//go:generate /usr/bin/env python ./misc/codegen.py

package curl
/*
#include <curl/curl.h>
#include "compat.h"
*/
import "C"

// CURLcode
const (
	E_OK                      = C.CURLE_OK
	E_UNSUPPORTED_PROTOCOL    = C.CURLE_UNSUPPORTED_PROTOCOL
	E_FAILED_INIT             = C.CURLE_FAILED_INIT
	E_URL_MALFORMAT           = C.CURLE_URL_MALFORMAT
	E_NOT_BUILT_IN            = C.CURLE_NOT_BUILT_IN
	E_COULDNT_RESOLVE_PROXY   = C.CURLE_COULDNT_RESOLVE_PROXY
	E_COULDNT_RESOLVE_HOST    = C.CURLE_COULDNT_RESOLVE_HOST
	E_COULDNT_CONNECT         = C.CURLE_COULDNT_CONNECT
	E_WEIRD_SERVER_REPLY      = C.CURLE_WEIRD_SERVER_REPLY
	E_REMOTE_ACCESS_DENIED    = C.CURLE_REMOTE_ACCESS_DENIED
	E_FTP_ACCEPT_FAILED       = C.CURLE_FTP_ACCEPT_FAILED
	E_FTP_WEIRD_PASS_REPLY    = C.CURLE_FTP_WEIRD_PASS_REPLY
	E_FTP_ACCEPT_TIMEOUT      = C.CURLE_FTP_ACCEPT_TIMEOUT
	E_FTP_WEIRD_PASV_REPLY    = C.CURLE_FTP_WEIRD_PASV_REPLY
	E_FTP_WEIRD_227_FORMAT    = C.CURLE_FTP_WEIRD_227_FORMAT
	E_FTP_CANT_GET_HOST       = C.CURLE_FTP_CANT_GET_HOST
	E_HTTP2                   = C.CURLE_HTTP2
	E_FTP_COULDNT_SET_TYPE    = C.CURLE_FTP_COULDNT_SET_TYPE
	E_PARTIAL_FILE            = C.CURLE_PARTIAL_FILE
	E_FTP_COULDNT_RETR_FILE   = C.CURLE_FTP_COULDNT_RETR_FILE
	E_OBSOLETE20              = C.CURLE_OBSOLETE20
	E_QUOTE_ERROR             = C.CURLE_QUOTE_ERROR
	E_HTTP_RETURNED_ERROR     = C.CURLE_HTTP_RETURNED_ERROR
	E_WRITE_ERROR             = C.CURLE_WRITE_ERROR
	E_OBSOLETE24              = C.CURLE_OBSOLETE24
	E_UPLOAD_FAILED           = C.CURLE_UPLOAD_FAILED
	E_READ_ERROR              = C.CURLE_READ_ERROR
	E_OUT_OF_MEMORY           = C.CURLE_OUT_OF_MEMORY
	E_OPERATION_TIMEDOUT      = C.CURLE_OPERATION_TIMEDOUT
	E_OBSOLETE29              = C.CURLE_OBSOLETE29
	E_FTP_PORT_FAILED         = C.CURLE_FTP_PORT_FAILED
	E_FTP_COULDNT_USE_REST    = C.CURLE_FTP_COULDNT_USE_REST
	E_OBSOLETE32              = C.CURLE_OBSOLETE32
	E_RANGE_ERROR             = C.CURLE_RANGE_ERROR
	E_HTTP_POST_ERROR         = C.CURLE_HTTP_POST_ERROR
	E_SSL_CONNECT_ERROR       = C.CURLE_SSL_CONNECT_ERROR
	E_BAD_DOWNLOAD_RESUME     = C.CURLE_BAD_DOWNLOAD_RESUME
	E_FILE_COULDNT_READ_FILE  = C.CURLE_FILE_COULDNT_READ_FILE
	E_LDAP_CANNOT_BIND        = C.CURLE_LDAP_CANNOT_BIND
	E_LDAP_SEARCH_FAILED      = C.CURLE_LDAP_SEARCH_FAILED
	E_OBSOLETE40              = C.CURLE_OBSOLETE40
	E_FUNCTION_NOT_FOUND      = C.CURLE_FUNCTION_NOT_FOUND
	E_ABORTED_BY_CALLBACK     = C.CURLE_ABORTED_BY_CALLBACK
	E_BAD_FUNCTION_ARGUMENT   = C.CURLE_BAD_FUNCTION_ARGUMENT
	E_OBSOLETE44              = C.CURLE_OBSOLETE44
	E_INTERFACE_FAILED        = C.CURLE_INTERFACE_FAILED
	E_OBSOLETE46              = C.CURLE_OBSOLETE46
	E_TOO_MANY_REDIRECTS      = C.CURLE_TOO_MANY_REDIRECTS
	E_UNKNOWN_OPTION          = C.CURLE_UNKNOWN_OPTION
	E_SETOPT_OPTION_SYNTAX    = C.CURLE_SETOPT_OPTION_SYNTAX
	E_OBSOLETE50              = C.CURLE_OBSOLETE50
	E_OBSOLETE51              = C.CURLE_OBSOLETE51
	E_GOT_NOTHING             = C.CURLE_GOT_NOTHING
	E_SSL_ENGINE_NOTFOUND     = C.CURLE_SSL_ENGINE_NOTFOUND
	E_SSL_ENGINE_SETFAILED    = C.CURLE_SSL_ENGINE_SETFAILED
	E_SEND_ERROR              = C.CURLE_SEND_ERROR
	E_RECV_ERROR              = C.CURLE_RECV_ERROR
	E_OBSOLETE57              = C.CURLE_OBSOLETE57
	E_SSL_CERTPROBLEM         = C.CURLE_SSL_CERTPROBLEM
	E_SSL_CIPHER              = C.CURLE_SSL_CIPHER
	E_PEER_FAILED_VERIFICATION = C.CURLE_PEER_FAILED_VERIFICATION
	E_BAD_CONTENT_ENCODING    = C.CURLE_BAD_CONTENT_ENCODING
	E_OBSOLETE62              = C.CURLE_OBSOLETE62
	E_FILESIZE_EXCEEDED       = C.CURLE_FILESIZE_EXCEEDED
	E_USE_SSL_FAILED          = C.CURLE_USE_SSL_FAILED
	E_SEND_FAIL_REWIND        = C.CURLE_SEND_FAIL_REWIND
	E_SSL_ENGINE_INITFAILED   = C.CURLE_SSL_ENGINE_INITFAILED
	E_LOGIN_DENIED            = C.CURLE_LOGIN_DENIED
	E_TFTP_NOTFOUND           = C.CURLE_TFTP_NOTFOUND
	E_TFTP_PERM               = C.CURLE_TFTP_PERM
	E_REMOTE_DISK_FULL        = C.CURLE_REMOTE_DISK_FULL
	E_TFTP_ILLEGAL            = C.CURLE_TFTP_ILLEGAL
	E_TFTP_UNKNOWNID          = C.CURLE_TFTP_UNKNOWNID
	E_REMOTE_FILE_EXISTS      = C.CURLE_REMOTE_FILE_EXISTS
	E_TFTP_NOSUCHUSER         = C.CURLE_TFTP_NOSUCHUSER
	E_OBSOLETE75              = C.CURLE_OBSOLETE75
	E_OBSOLETE76              = C.CURLE_OBSOLETE76
	E_SSL_CACERT_BADFILE      = C.CURLE_SSL_CACERT_BADFILE
	E_REMOTE_FILE_NOT_FOUND   = C.CURLE_REMOTE_FILE_NOT_FOUND
	E_SSH                     = C.CURLE_SSH
	E_SSL_SHUTDOWN_FAILED     = C.CURLE_SSL_SHUTDOWN_FAILED
	E_AGAIN                   = C.CURLE_AGAIN
	E_SSL_CRL_BADFILE         = C.CURLE_SSL_CRL_BADFILE
	E_SSL_ISSUER_ERROR        = C.CURLE_SSL_ISSUER_ERROR
	E_FTP_PRET_FAILED         = C.CURLE_FTP_PRET_FAILED
	E_RTSP_CSEQ_ERROR         = C.CURLE_RTSP_CSEQ_ERROR
	E_RTSP_SESSION_ERROR      = C.CURLE_RTSP_SESSION_ERROR
	E_FTP_BAD_FILE_LIST       = C.CURLE_FTP_BAD_FILE_LIST
	E_CHUNK_FAILED            = C.CURLE_CHUNK_FAILED
	E_NO_CONNECTION_AVAILABLE = C.CURLE_NO_CONNECTION_AVAILABLE
	E_SSL_PINNEDPUBKEYNOTMATCH = C.CURLE_SSL_PINNEDPUBKEYNOTMATCH
	E_SSL_INVALIDCERTSTATUS   = C.CURLE_SSL_INVALIDCERTSTATUS
	E_HTTP2_STREAM            = C.CURLE_HTTP2_STREAM
	E_RECURSIVE_API_CALL      = C.CURLE_RECURSIVE_API_CALL
	E_AUTH_ERROR              = C.CURLE_AUTH_ERROR
	E_HTTP3                   = C.CURLE_HTTP3
	E_QUIC_CONNECT_ERROR      = C.CURLE_QUIC_CONNECT_ERROR
	E_PROXY                   = C.CURLE_PROXY
	E_SSL_CLIENTCERT          = C.CURLE_SSL_CLIENTCERT
	E_UNRECOVERABLE_POLL      = C.CURLE_UNRECOVERABLE_POLL
	E_TOO_LARGE               = C.CURLE_TOO_LARGE
	E_ECH_REQUIRED            = C.CURLE_ECH_REQUIRED
	E_OBSOLETE16              = C.CURLE_OBSOLETE16
	E_OBSOLETE10              = C.CURLE_OBSOLETE10
	E_OBSOLETE12              = C.CURLE_OBSOLETE12
	E_FTP_WEIRD_SERVER_REPLY  = C.CURLE_FTP_WEIRD_SERVER_REPLY
	E_SSL_CACERT              = C.CURLE_SSL_CACERT
	E_UNKNOWN_TELNET_OPTION   = C.CURLE_UNKNOWN_TELNET_OPTION
	E_TELNET_OPTION_SYNTAX    = C.CURLE_TELNET_OPTION_SYNTAX
	E_SSL_PEER_CERTIFICATE    = C.CURLE_SSL_PEER_CERTIFICATE
	E_OBSOLETE                = C.CURLE_OBSOLETE
	E_BAD_PASSWORD_ENTERED    = C.CURLE_BAD_PASSWORD_ENTERED
	E_BAD_CALLING_ORDER       = C.CURLE_BAD_CALLING_ORDER
	E_FTP_USER_PASSWORD_INCORRECT = C.CURLE_FTP_USER_PASSWORD_INCORRECT
	E_FTP_CANT_RECONNECT      = C.CURLE_FTP_CANT_RECONNECT
	E_FTP_COULDNT_GET_SIZE    = C.CURLE_FTP_COULDNT_GET_SIZE
	E_FTP_COULDNT_SET_ASCII   = C.CURLE_FTP_COULDNT_SET_ASCII
	E_FTP_WEIRD_USER_REPLY    = C.CURLE_FTP_WEIRD_USER_REPLY
	E_FTP_WRITE_ERROR         = C.CURLE_FTP_WRITE_ERROR
	E_LIBRARY_NOT_FOUND       = C.CURLE_LIBRARY_NOT_FOUND
	E_MALFORMAT_USER          = C.CURLE_MALFORMAT_USER
	E_SHARE_IN_USE            = C.CURLE_SHARE_IN_USE
	E_URL_MALFORMAT_USER      = C.CURLE_URL_MALFORMAT_USER
	E_FTP_ACCESS_DENIED       = C.CURLE_FTP_ACCESS_DENIED
	E_FTP_COULDNT_SET_BINARY  = C.CURLE_FTP_COULDNT_SET_BINARY
	E_FTP_QUOTE_ERROR         = C.CURLE_FTP_QUOTE_ERROR
	E_TFTP_DISKFULL           = C.CURLE_TFTP_DISKFULL
	E_TFTP_EXISTS             = C.CURLE_TFTP_EXISTS
	E_HTTP_RANGE_ERROR        = C.CURLE_HTTP_RANGE_ERROR
	E_FTP_SSL_FAILED          = C.CURLE_FTP_SSL_FAILED
	E_OPERATION_TIMEOUTED     = C.CURLE_OPERATION_TIMEOUTED
	E_HTTP_NOT_FOUND          = C.CURLE_HTTP_NOT_FOUND
	E_HTTP_PORT_FAILED        = C.CURLE_HTTP_PORT_FAILED
	E_FTP_COULDNT_STOR_FILE   = C.CURLE_FTP_COULDNT_STOR_FILE
	E_FTP_PARTIAL_FILE        = C.CURLE_FTP_PARTIAL_FILE
	E_FTP_BAD_DOWNLOAD_RESUME = C.CURLE_FTP_BAD_DOWNLOAD_RESUME
	E_LDAP_INVALID_URL        = C.CURLE_LDAP_INVALID_URL
	E_CONV_REQD               = C.CURLE_CONV_REQD
	E_CONV_FAILED             = C.CURLE_CONV_FAILED
	E_ALREADY_COMPLETE        = C.CURLE_ALREADY_COMPLETE
)

// easy.Setopt(flag, ...)
const (
	OPT_ENCODING                  = C.CURLOPT_ENCODING
	OPT_FILE                      = C.CURLOPT_FILE
	OPT_INFILE                    = C.CURLOPT_INFILE
	OPT_WRITEHEADER               = C.CURLOPT_WRITEHEADER
	OPT_WRITEINFO                 = C.CURLOPT_WRITEINFO
	OPT_CLOSEPOLICY               = C.CURLOPT_CLOSEPOLICY
	OPT_WRITEDATA                 = C.CURLOPT_WRITEDATA
	OPT_URL                       = C.CURLOPT_URL
	OPT_PORT                      = C.CURLOPT_PORT
	OPT_PROXY                     = C.CURLOPT_PROXY
	OPT_USERPWD                   = C.CURLOPT_USERPWD
	OPT_PROXYUSERPWD              = C.CURLOPT_PROXYUSERPWD
	OPT_RANGE                     = C.CURLOPT_RANGE
	OPT_READDATA                  = C.CURLOPT_READDATA
	OPT_ERRORBUFFER               = C.CURLOPT_ERRORBUFFER
	OPT_WRITEFUNCTION             = C.CURLOPT_WRITEFUNCTION
	OPT_READFUNCTION              = C.CURLOPT_READFUNCTION
	OPT_TIMEOUT                   = C.CURLOPT_TIMEOUT
	OPT_INFILESIZE                = C.CURLOPT_INFILESIZE
	OPT_POSTFIELDS                = C.CURLOPT_POSTFIELDS
	OPT_REFERER                   = C.CURLOPT_REFERER
	OPT_FTPPORT                   = C.CURLOPT_FTPPORT
	OPT_USERAGENT                 = C.CURLOPT_USERAGENT
	OPT_LOW_SPEED_LIMIT           = C.CURLOPT_LOW_SPEED_LIMIT
	OPT_LOW_SPEED_TIME            = C.CURLOPT_LOW_SPEED_TIME
	OPT_RESUME_FROM               = C.CURLOPT_RESUME_FROM
	OPT_COOKIE                    = C.CURLOPT_COOKIE
	OPT_HTTPHEADER                = C.CURLOPT_HTTPHEADER
	OPT_HTTPPOST                  = C.CURLOPT_HTTPPOST
	OPT_SSLCERT                   = C.CURLOPT_SSLCERT
	OPT_KEYPASSWD                 = C.CURLOPT_KEYPASSWD
	OPT_CRLF                      = C.CURLOPT_CRLF
	OPT_QUOTE                     = C.CURLOPT_QUOTE
	OPT_HEADERDATA                = C.CURLOPT_HEADERDATA
	OPT_COOKIEFILE                = C.CURLOPT_COOKIEFILE
	OPT_SSLVERSION                = C.CURLOPT_SSLVERSION
	OPT_TIMECONDITION             = C.CURLOPT_TIMECONDITION
	OPT_TIMEVALUE                 = C.CURLOPT_TIMEVALUE
	OPT_CUSTOMREQUEST             = C.CURLOPT_CUSTOMREQUEST
	OPT_STDERR                    = C.CURLOPT_STDERR
	OPT_POSTQUOTE                 = C.CURLOPT_POSTQUOTE
	OPT_OBSOLETE40                = C.CURLOPT_OBSOLETE40
	OPT_VERBOSE                   = C.CURLOPT_VERBOSE
	OPT_HEADER                    = C.CURLOPT_HEADER
	OPT_NOPROGRESS                = C.CURLOPT_NOPROGRESS
	OPT_NOBODY                    = C.CURLOPT_NOBODY
	OPT_FAILONERROR               = C.CURLOPT_FAILONERROR
	OPT_UPLOAD                    = C.CURLOPT_UPLOAD
	OPT_POST                      = C.CURLOPT_POST
	OPT_DIRLISTONLY               = C.CURLOPT_DIRLISTONLY
	OPT_APPEND                    = C.CURLOPT_APPEND
	OPT_NETRC                     = C.CURLOPT_NETRC
	OPT_FOLLOWLOCATION            = C.CURLOPT_FOLLOWLOCATION
	OPT_TRANSFERTEXT              = C.CURLOPT_TRANSFERTEXT
	OPT_PUT                       = C.CURLOPT_PUT
	OPT_PROGRESSFUNCTION          = C.CURLOPT_PROGRESSFUNCTION
	OPT_XFERINFODATA              = C.CURLOPT_XFERINFODATA
	OPT_PROGRESSDATA              = C.CURLOPT_PROGRESSDATA
	OPT_AUTOREFERER               = C.CURLOPT_AUTOREFERER
	OPT_PROXYPORT                 = C.CURLOPT_PROXYPORT
	OPT_POSTFIELDSIZE             = C.CURLOPT_POSTFIELDSIZE
	OPT_HTTPPROXYTUNNEL           = C.CURLOPT_HTTPPROXYTUNNEL
	OPT_INTERFACE                 = C.CURLOPT_INTERFACE
	OPT_KRBLEVEL                  = C.CURLOPT_KRBLEVEL
	OPT_SSL_VERIFYPEER            = C.CURLOPT_SSL_VERIFYPEER
	OPT_CAINFO                    = C.CURLOPT_CAINFO
	OPT_MAXREDIRS                 = C.CURLOPT_MAXREDIRS
	OPT_FILETIME                  = C.CURLOPT_FILETIME
	OPT_TELNETOPTIONS             = C.CURLOPT_TELNETOPTIONS
	OPT_MAXCONNECTS               = C.CURLOPT_MAXCONNECTS
	OPT_OBSOLETE72                = C.CURLOPT_OBSOLETE72
	OPT_FRESH_CONNECT             = C.CURLOPT_FRESH_CONNECT
	OPT_FORBID_REUSE              = C.CURLOPT_FORBID_REUSE
	OPT_RANDOM_FILE               = C.CURLOPT_RANDOM_FILE
	OPT_EGDSOCKET                 = C.CURLOPT_EGDSOCKET
	OPT_CONNECTTIMEOUT            = C.CURLOPT_CONNECTTIMEOUT
	OPT_HEADERFUNCTION            = C.CURLOPT_HEADERFUNCTION
	OPT_HTTPGET                   = C.CURLOPT_HTTPGET
	OPT_SSL_VERIFYHOST            = C.CURLOPT_SSL_VERIFYHOST
	OPT_COOKIEJAR                 = C.CURLOPT_COOKIEJAR
	OPT_SSL_CIPHER_LIST           = C.CURLOPT_SSL_CIPHER_LIST
	OPT_HTTP_VERSION              = C.CURLOPT_HTTP_VERSION
	OPT_FTP_USE_EPSV              = C.CURLOPT_FTP_USE_EPSV
	OPT_SSLCERTTYPE               = C.CURLOPT_SSLCERTTYPE
	OPT_SSLKEY                    = C.CURLOPT_SSLKEY
	OPT_SSLKEYTYPE                = C.CURLOPT_SSLKEYTYPE
	OPT_SSLENGINE                 = C.CURLOPT_SSLENGINE
	OPT_SSLENGINE_DEFAULT         = C.CURLOPT_SSLENGINE_DEFAULT
	OPT_DNS_USE_GLOBAL_CACHE      = C.CURLOPT_DNS_USE_GLOBAL_CACHE
	OPT_DNS_CACHE_TIMEOUT         = C.CURLOPT_DNS_CACHE_TIMEOUT
	OPT_PREQUOTE                  = C.CURLOPT_PREQUOTE
	OPT_DEBUGFUNCTION             = C.CURLOPT_DEBUGFUNCTION
	OPT_DEBUGDATA                 = C.CURLOPT_DEBUGDATA
	OPT_COOKIESESSION             = C.CURLOPT_COOKIESESSION
	OPT_CAPATH                    = C.CURLOPT_CAPATH
	OPT_BUFFERSIZE                = C.CURLOPT_BUFFERSIZE
	OPT_NOSIGNAL                  = C.CURLOPT_NOSIGNAL
	OPT_SHARE                     = C.CURLOPT_SHARE
	OPT_PROXYTYPE                 = C.CURLOPT_PROXYTYPE
	OPT_ACCEPT_ENCODING           = C.CURLOPT_ACCEPT_ENCODING
	OPT_PRIVATE                   = C.CURLOPT_PRIVATE
	OPT_HTTP200ALIASES            = C.CURLOPT_HTTP200ALIASES
	OPT_UNRESTRICTED_AUTH         = C.CURLOPT_UNRESTRICTED_AUTH
	OPT_FTP_USE_EPRT              = C.CURLOPT_FTP_USE_EPRT
	OPT_HTTPAUTH                  = C.CURLOPT_HTTPAUTH
	OPT_SSL_CTX_FUNCTION          = C.CURLOPT_SSL_CTX_FUNCTION
	OPT_SSL_CTX_DATA              = C.CURLOPT_SSL_CTX_DATA
	OPT_FTP_CREATE_MISSING_DIRS   = C.CURLOPT_FTP_CREATE_MISSING_DIRS
	OPT_PROXYAUTH                 = C.CURLOPT_PROXYAUTH
	OPT_SERVER_RESPONSE_TIMEOUT   = C.CURLOPT_SERVER_RESPONSE_TIMEOUT
	OPT_IPRESOLVE                 = C.CURLOPT_IPRESOLVE
	OPT_MAXFILESIZE               = C.CURLOPT_MAXFILESIZE
	OPT_INFILESIZE_LARGE          = C.CURLOPT_INFILESIZE_LARGE
	OPT_RESUME_FROM_LARGE         = C.CURLOPT_RESUME_FROM_LARGE
	OPT_MAXFILESIZE_LARGE         = C.CURLOPT_MAXFILESIZE_LARGE
	OPT_NETRC_FILE                = C.CURLOPT_NETRC_FILE
	OPT_USE_SSL                   = C.CURLOPT_USE_SSL
	OPT_POSTFIELDSIZE_LARGE       = C.CURLOPT_POSTFIELDSIZE_LARGE
	OPT_TCP_NODELAY               = C.CURLOPT_TCP_NODELAY
	OPT_FTPSSLAUTH                = C.CURLOPT_FTPSSLAUTH
	OPT_IOCTLFUNCTION             = C.CURLOPT_IOCTLFUNCTION
	OPT_IOCTLDATA                 = C.CURLOPT_IOCTLDATA
	OPT_FTP_ACCOUNT               = C.CURLOPT_FTP_ACCOUNT
	OPT_COOKIELIST                = C.CURLOPT_COOKIELIST
	OPT_IGNORE_CONTENT_LENGTH     = C.CURLOPT_IGNORE_CONTENT_LENGTH
	OPT_FTP_SKIP_PASV_IP          = C.CURLOPT_FTP_SKIP_PASV_IP
	OPT_FTP_FILEMETHOD            = C.CURLOPT_FTP_FILEMETHOD
	OPT_LOCALPORT                 = C.CURLOPT_LOCALPORT
	OPT_LOCALPORTRANGE            = C.CURLOPT_LOCALPORTRANGE
	OPT_CONNECT_ONLY              = C.CURLOPT_CONNECT_ONLY
	OPT_CONV_FROM_NETWORK_FUNCTION = C.CURLOPT_CONV_FROM_NETWORK_FUNCTION
	OPT_CONV_TO_NETWORK_FUNCTION  = C.CURLOPT_CONV_TO_NETWORK_FUNCTION
	OPT_CONV_FROM_UTF8_FUNCTION   = C.CURLOPT_CONV_FROM_UTF8_FUNCTION
	OPT_MAX_SEND_SPEED_LARGE      = C.CURLOPT_MAX_SEND_SPEED_LARGE
	OPT_MAX_RECV_SPEED_LARGE      = C.CURLOPT_MAX_RECV_SPEED_LARGE
	OPT_FTP_ALTERNATIVE_TO_USER   = C.CURLOPT_FTP_ALTERNATIVE_TO_USER
	OPT_SOCKOPTFUNCTION           = C.CURLOPT_SOCKOPTFUNCTION
	OPT_SOCKOPTDATA               = C.CURLOPT_SOCKOPTDATA
	OPT_SSL_SESSIONID_CACHE       = C.CURLOPT_SSL_SESSIONID_CACHE
	OPT_SSH_AUTH_TYPES            = C.CURLOPT_SSH_AUTH_TYPES
	OPT_SSH_PUBLIC_KEYFILE        = C.CURLOPT_SSH_PUBLIC_KEYFILE
	OPT_SSH_PRIVATE_KEYFILE       = C.CURLOPT_SSH_PRIVATE_KEYFILE
	OPT_FTP_SSL_CCC               = C.CURLOPT_FTP_SSL_CCC
	OPT_TIMEOUT_MS                = C.CURLOPT_TIMEOUT_MS
	OPT_CONNECTTIMEOUT_MS         = C.CURLOPT_CONNECTTIMEOUT_MS
	OPT_HTTP_TRANSFER_DECODING    = C.CURLOPT_HTTP_TRANSFER_DECODING
	OPT_HTTP_CONTENT_DECODING     = C.CURLOPT_HTTP_CONTENT_DECODING
	OPT_NEW_FILE_PERMS            = C.CURLOPT_NEW_FILE_PERMS
	OPT_NEW_DIRECTORY_PERMS       = C.CURLOPT_NEW_DIRECTORY_PERMS
	OPT_POSTREDIR                 = C.CURLOPT_POSTREDIR
	OPT_SSH_HOST_PUBLIC_KEY_MD5   = C.CURLOPT_SSH_HOST_PUBLIC_KEY_MD5
	OPT_OPENSOCKETFUNCTION        = C.CURLOPT_OPENSOCKETFUNCTION
	OPT_OPENSOCKETDATA            = C.CURLOPT_OPENSOCKETDATA
	OPT_COPYPOSTFIELDS            = C.CURLOPT_COPYPOSTFIELDS
	OPT_PROXY_TRANSFER_MODE       = C.CURLOPT_PROXY_TRANSFER_MODE
	OPT_SEEKFUNCTION              = C.CURLOPT_SEEKFUNCTION
	OPT_SEEKDATA                  = C.CURLOPT_SEEKDATA
	OPT_CRLFILE                   = C.CURLOPT_CRLFILE
	OPT_ISSUERCERT                = C.CURLOPT_ISSUERCERT
	OPT_ADDRESS_SCOPE             = C.CURLOPT_ADDRESS_SCOPE
	OPT_CERTINFO                  = C.CURLOPT_CERTINFO
	OPT_USERNAME                  = C.CURLOPT_USERNAME
	OPT_PASSWORD                  = C.CURLOPT_PASSWORD
	OPT_PROXYUSERNAME             = C.CURLOPT_PROXYUSERNAME
	OPT_PROXYPASSWORD             = C.CURLOPT_PROXYPASSWORD
	OPT_NOPROXY                   = C.CURLOPT_NOPROXY
	OPT_TFTP_BLKSIZE              = C.CURLOPT_TFTP_BLKSIZE
	OPT_SOCKS5_GSSAPI_SERVICE     = C.CURLOPT_SOCKS5_GSSAPI_SERVICE
	OPT_SOCKS5_GSSAPI_NEC         = C.CURLOPT_SOCKS5_GSSAPI_NEC
	OPT_PROTOCOLS                 = C.CURLOPT_PROTOCOLS
	OPT_REDIR_PROTOCOLS           = C.CURLOPT_REDIR_PROTOCOLS
	OPT_SSH_KNOWNHOSTS            = C.CURLOPT_SSH_KNOWNHOSTS
	OPT_SSH_KEYFUNCTION           = C.CURLOPT_SSH_KEYFUNCTION
	OPT_SSH_KEYDATA               = C.CURLOPT_SSH_KEYDATA
	OPT_MAIL_FROM                 = C.CURLOPT_MAIL_FROM
	OPT_MAIL_RCPT                 = C.CURLOPT_MAIL_RCPT
	OPT_FTP_USE_PRET              = C.CURLOPT_FTP_USE_PRET
	OPT_RTSP_REQUEST              = C.CURLOPT_RTSP_REQUEST
	OPT_RTSP_SESSION_ID           = C.CURLOPT_RTSP_SESSION_ID
	OPT_RTSP_STREAM_URI           = C.CURLOPT_RTSP_STREAM_URI
	OPT_RTSP_TRANSPORT            = C.CURLOPT_RTSP_TRANSPORT
	OPT_RTSP_CLIENT_CSEQ          = C.CURLOPT_RTSP_CLIENT_CSEQ
	OPT_RTSP_SERVER_CSEQ          = C.CURLOPT_RTSP_SERVER_CSEQ
	OPT_INTERLEAVEDATA            = C.CURLOPT_INTERLEAVEDATA
	OPT_INTERLEAVEFUNCTION        = C.CURLOPT_INTERLEAVEFUNCTION
	OPT_WILDCARDMATCH             = C.CURLOPT_WILDCARDMATCH
	OPT_CHUNK_BGN_FUNCTION        = C.CURLOPT_CHUNK_BGN_FUNCTION
	OPT_CHUNK_END_FUNCTION        = C.CURLOPT_CHUNK_END_FUNCTION
	OPT_FNMATCH_FUNCTION          = C.CURLOPT_FNMATCH_FUNCTION
	OPT_CHUNK_DATA                = C.CURLOPT_CHUNK_DATA
	OPT_FNMATCH_DATA              = C.CURLOPT_FNMATCH_DATA
	OPT_RESOLVE                   = C.CURLOPT_RESOLVE
	OPT_TLSAUTH_USERNAME          = C.CURLOPT_TLSAUTH_USERNAME
	OPT_TLSAUTH_PASSWORD          = C.CURLOPT_TLSAUTH_PASSWORD
	OPT_TLSAUTH_TYPE              = C.CURLOPT_TLSAUTH_TYPE
	OPT_TRANSFER_ENCODING         = C.CURLOPT_TRANSFER_ENCODING
	OPT_CLOSESOCKETFUNCTION       = C.CURLOPT_CLOSESOCKETFUNCTION
	OPT_CLOSESOCKETDATA           = C.CURLOPT_CLOSESOCKETDATA
	OPT_GSSAPI_DELEGATION         = C.CURLOPT_GSSAPI_DELEGATION
	OPT_DNS_SERVERS               = C.CURLOPT_DNS_SERVERS
	OPT_ACCEPTTIMEOUT_MS          = C.CURLOPT_ACCEPTTIMEOUT_MS
	OPT_TCP_KEEPALIVE             = C.CURLOPT_TCP_KEEPALIVE
	OPT_TCP_KEEPIDLE              = C.CURLOPT_TCP_KEEPIDLE
	OPT_TCP_KEEPINTVL             = C.CURLOPT_TCP_KEEPINTVL
	OPT_SSL_OPTIONS               = C.CURLOPT_SSL_OPTIONS
	OPT_MAIL_AUTH                 = C.CURLOPT_MAIL_AUTH
	OPT_SASL_IR                   = C.CURLOPT_SASL_IR
	OPT_XFERINFOFUNCTION          = C.CURLOPT_XFERINFOFUNCTION
	OPT_XOAUTH2_BEARER            = C.CURLOPT_XOAUTH2_BEARER
	OPT_DNS_INTERFACE             = C.CURLOPT_DNS_INTERFACE
	OPT_DNS_LOCAL_IP4             = C.CURLOPT_DNS_LOCAL_IP4
	OPT_DNS_LOCAL_IP6             = C.CURLOPT_DNS_LOCAL_IP6
	OPT_LOGIN_OPTIONS             = C.CURLOPT_LOGIN_OPTIONS
	OPT_SSL_ENABLE_NPN            = C.CURLOPT_SSL_ENABLE_NPN
	OPT_SSL_ENABLE_ALPN           = C.CURLOPT_SSL_ENABLE_ALPN
	OPT_EXPECT_100_TIMEOUT_MS     = C.CURLOPT_EXPECT_100_TIMEOUT_MS
	OPT_PROXYHEADER               = C.CURLOPT_PROXYHEADER
	OPT_HEADEROPT                 = C.CURLOPT_HEADEROPT
	OPT_PINNEDPUBLICKEY           = C.CURLOPT_PINNEDPUBLICKEY
	OPT_UNIX_SOCKET_PATH          = C.CURLOPT_UNIX_SOCKET_PATH
	OPT_SSL_VERIFYSTATUS          = C.CURLOPT_SSL_VERIFYSTATUS
	OPT_SSL_FALSESTART            = C.CURLOPT_SSL_FALSESTART
	OPT_PATH_AS_IS                = C.CURLOPT_PATH_AS_IS
	OPT_PROXY_SERVICE_NAME        = C.CURLOPT_PROXY_SERVICE_NAME
	OPT_SERVICE_NAME              = C.CURLOPT_SERVICE_NAME
	OPT_PIPEWAIT                  = C.CURLOPT_PIPEWAIT
	OPT_DEFAULT_PROTOCOL          = C.CURLOPT_DEFAULT_PROTOCOL
	OPT_STREAM_WEIGHT             = C.CURLOPT_STREAM_WEIGHT
	OPT_STREAM_DEPENDS            = C.CURLOPT_STREAM_DEPENDS
	OPT_STREAM_DEPENDS_E          = C.CURLOPT_STREAM_DEPENDS_E
	OPT_TFTP_NO_OPTIONS           = C.CURLOPT_TFTP_NO_OPTIONS
	OPT_CONNECT_TO                = C.CURLOPT_CONNECT_TO
	OPT_TCP_FASTOPEN              = C.CURLOPT_TCP_FASTOPEN
	OPT_KEEP_SENDING_ON_ERROR     = C.CURLOPT_KEEP_SENDING_ON_ERROR
	OPT_PROXY_CAINFO              = C.CURLOPT_PROXY_CAINFO
	OPT_PROXY_CAPATH              = C.CURLOPT_PROXY_CAPATH
	OPT_PROXY_SSL_VERIFYPEER      = C.CURLOPT_PROXY_SSL_VERIFYPEER
	OPT_PROXY_SSL_VERIFYHOST      = C.CURLOPT_PROXY_SSL_VERIFYHOST
	OPT_PROXY_SSLVERSION          = C.CURLOPT_PROXY_SSLVERSION
	OPT_PROXY_TLSAUTH_USERNAME    = C.CURLOPT_PROXY_TLSAUTH_USERNAME
	OPT_PROXY_TLSAUTH_PASSWORD    = C.CURLOPT_PROXY_TLSAUTH_PASSWORD
	OPT_PROXY_TLSAUTH_TYPE        = C.CURLOPT_PROXY_TLSAUTH_TYPE
	OPT_PROXY_SSLCERT             = C.CURLOPT_PROXY_SSLCERT
	OPT_PROXY_SSLCERTTYPE         = C.CURLOPT_PROXY_SSLCERTTYPE
	OPT_PROXY_SSLKEY              = C.CURLOPT_PROXY_SSLKEY
	OPT_PROXY_SSLKEYTYPE          = C.CURLOPT_PROXY_SSLKEYTYPE
	OPT_PROXY_KEYPASSWD           = C.CURLOPT_PROXY_KEYPASSWD
	OPT_PROXY_SSL_CIPHER_LIST     = C.CURLOPT_PROXY_SSL_CIPHER_LIST
	OPT_PROXY_CRLFILE             = C.CURLOPT_PROXY_CRLFILE
	OPT_PROXY_SSL_OPTIONS         = C.CURLOPT_PROXY_SSL_OPTIONS
	OPT_PRE_PROXY                 = C.CURLOPT_PRE_PROXY
	OPT_PROXY_PINNEDPUBLICKEY     = C.CURLOPT_PROXY_PINNEDPUBLICKEY
	OPT_ABSTRACT_UNIX_SOCKET      = C.CURLOPT_ABSTRACT_UNIX_SOCKET
	OPT_SUPPRESS_CONNECT_HEADERS  = C.CURLOPT_SUPPRESS_CONNECT_HEADERS
	OPT_REQUEST_TARGET            = C.CURLOPT_REQUEST_TARGET
	OPT_SOCKS5_AUTH               = C.CURLOPT_SOCKS5_AUTH
	OPT_SSH_COMPRESSION           = C.CURLOPT_SSH_COMPRESSION
	OPT_MIMEPOST                  = C.CURLOPT_MIMEPOST
	OPT_TIMEVALUE_LARGE           = C.CURLOPT_TIMEVALUE_LARGE
	OPT_HAPPY_EYEBALLS_TIMEOUT_MS = C.CURLOPT_HAPPY_EYEBALLS_TIMEOUT_MS
	OPT_RESOLVER_START_FUNCTION   = C.CURLOPT_RESOLVER_START_FUNCTION
	OPT_RESOLVER_START_DATA       = C.CURLOPT_RESOLVER_START_DATA
	OPT_HAPROXYPROTOCOL           = C.CURLOPT_HAPROXYPROTOCOL
	OPT_DNS_SHUFFLE_ADDRESSES     = C.CURLOPT_DNS_SHUFFLE_ADDRESSES
	OPT_TLS13_CIPHERS             = C.CURLOPT_TLS13_CIPHERS
	OPT_PROXY_TLS13_CIPHERS       = C.CURLOPT_PROXY_TLS13_CIPHERS
	OPT_DISALLOW_USERNAME_IN_URL  = C.CURLOPT_DISALLOW_USERNAME_IN_URL
	OPT_DOH_URL                   = C.CURLOPT_DOH_URL
	OPT_UPLOAD_BUFFERSIZE         = C.CURLOPT_UPLOAD_BUFFERSIZE
	OPT_UPKEEP_INTERVAL_MS        = C.CURLOPT_UPKEEP_INTERVAL_MS
	OPT_CURLU                     = C.CURLOPT_CURLU
	OPT_TRAILERFUNCTION           = C.CURLOPT_TRAILERFUNCTION
	OPT_TRAILERDATA               = C.CURLOPT_TRAILERDATA
	OPT_HTTP09_ALLOWED            = C.CURLOPT_HTTP09_ALLOWED
	OPT_ALTSVC_CTRL               = C.CURLOPT_ALTSVC_CTRL
	OPT_ALTSVC                    = C.CURLOPT_ALTSVC
	OPT_MAXAGE_CONN               = C.CURLOPT_MAXAGE_CONN
	OPT_SASL_AUTHZID              = C.CURLOPT_SASL_AUTHZID
	OPT_MAIL_RCPT_ALLOWFAILS      = C.CURLOPT_MAIL_RCPT_ALLOWFAILS
	OPT_SSLCERT_BLOB              = C.CURLOPT_SSLCERT_BLOB
	OPT_SSLKEY_BLOB               = C.CURLOPT_SSLKEY_BLOB
	OPT_PROXY_SSLCERT_BLOB        = C.CURLOPT_PROXY_SSLCERT_BLOB
	OPT_PROXY_SSLKEY_BLOB         = C.CURLOPT_PROXY_SSLKEY_BLOB
	OPT_ISSUERCERT_BLOB           = C.CURLOPT_ISSUERCERT_BLOB
	OPT_PROXY_ISSUERCERT          = C.CURLOPT_PROXY_ISSUERCERT
	OPT_PROXY_ISSUERCERT_BLOB     = C.CURLOPT_PROXY_ISSUERCERT_BLOB
	OPT_SSL_EC_CURVES             = C.CURLOPT_SSL_EC_CURVES
	OPT_HSTS_CTRL                 = C.CURLOPT_HSTS_CTRL
	OPT_HSTS                      = C.CURLOPT_HSTS
	OPT_HSTSREADFUNCTION          = C.CURLOPT_HSTSREADFUNCTION
	OPT_HSTSREADDATA              = C.CURLOPT_HSTSREADDATA
	OPT_HSTSWRITEFUNCTION         = C.CURLOPT_HSTSWRITEFUNCTION
	OPT_HSTSWRITEDATA             = C.CURLOPT_HSTSWRITEDATA
	OPT_AWS_SIGV4                 = C.CURLOPT_AWS_SIGV4
	OPT_DOH_SSL_VERIFYPEER        = C.CURLOPT_DOH_SSL_VERIFYPEER
	OPT_DOH_SSL_VERIFYHOST        = C.CURLOPT_DOH_SSL_VERIFYHOST
	OPT_DOH_SSL_VERIFYSTATUS      = C.CURLOPT_DOH_SSL_VERIFYSTATUS
	OPT_CAINFO_BLOB               = C.CURLOPT_CAINFO_BLOB
	OPT_PROXY_CAINFO_BLOB         = C.CURLOPT_PROXY_CAINFO_BLOB
	OPT_SSH_HOST_PUBLIC_KEY_SHA256 = C.CURLOPT_SSH_HOST_PUBLIC_KEY_SHA256
	OPT_PREREQFUNCTION            = C.CURLOPT_PREREQFUNCTION
	OPT_PREREQDATA                = C.CURLOPT_PREREQDATA
	OPT_MAXLIFETIME_CONN          = C.CURLOPT_MAXLIFETIME_CONN
	OPT_MIME_OPTIONS              = C.CURLOPT_MIME_OPTIONS
	OPT_SSH_HOSTKEYFUNCTION       = C.CURLOPT_SSH_HOSTKEYFUNCTION
	OPT_SSH_HOSTKEYDATA           = C.CURLOPT_SSH_HOSTKEYDATA
	OPT_PROTOCOLS_STR             = C.CURLOPT_PROTOCOLS_STR
	OPT_REDIR_PROTOCOLS_STR       = C.CURLOPT_REDIR_PROTOCOLS_STR
	OPT_WS_OPTIONS                = C.CURLOPT_WS_OPTIONS
	OPT_CA_CACHE_TIMEOUT          = C.CURLOPT_CA_CACHE_TIMEOUT
	OPT_QUICK_EXIT                = C.CURLOPT_QUICK_EXIT
	OPT_HAPROXY_CLIENT_IP         = C.CURLOPT_HAPROXY_CLIENT_IP
	OPT_SERVER_RESPONSE_TIMEOUT_MS = C.CURLOPT_SERVER_RESPONSE_TIMEOUT_MS
	OPT_ECH                       = C.CURLOPT_ECH
	OPT_TCP_KEEPCNT               = C.CURLOPT_TCP_KEEPCNT
	OPT_POST301                   = C.CURLOPT_POST301
	OPT_SSLKEYPASSWD              = C.CURLOPT_SSLKEYPASSWD
	OPT_FTPAPPEND                 = C.CURLOPT_FTPAPPEND
	OPT_FTPLISTONLY               = C.CURLOPT_FTPLISTONLY
	OPT_FTP_SSL                   = C.CURLOPT_FTP_SSL
	OPT_SSLCERTPASSWD             = C.CURLOPT_SSLCERTPASSWD
	OPT_KRB4LEVEL                 = C.CURLOPT_KRB4LEVEL
	OPT_FTP_RESPONSE_TIMEOUT      = C.CURLOPT_FTP_RESPONSE_TIMEOUT
	OPT_MAIL_RCPT_ALLLOWFAILS     = C.CURLOPT_MAIL_RCPT_ALLLOWFAILS
	OPT_RTSPHEADER                = C.CURLOPT_RTSPHEADER
)

// easy.Getinfo(flag)
const (
	INFO_TEXT                 = C.CURLINFO_TEXT
	INFO_END                  = C.CURLINFO_END
	INFO_EFFECTIVE_URL        = C.CURLINFO_EFFECTIVE_URL
	INFO_RESPONSE_CODE        = C.CURLINFO_RESPONSE_CODE
	INFO_TOTAL_TIME           = C.CURLINFO_TOTAL_TIME
	INFO_NAMELOOKUP_TIME      = C.CURLINFO_NAMELOOKUP_TIME
	INFO_CONNECT_TIME         = C.CURLINFO_CONNECT_TIME
	INFO_PRETRANSFER_TIME     = C.CURLINFO_PRETRANSFER_TIME
	INFO_SIZE_UPLOAD          = C.CURLINFO_SIZE_UPLOAD
	INFO_SIZE_UPLOAD_T        = C.CURLINFO_SIZE_UPLOAD_T
	INFO_SIZE_DOWNLOAD        = C.CURLINFO_SIZE_DOWNLOAD
	INFO_SIZE_DOWNLOAD_T      = C.CURLINFO_SIZE_DOWNLOAD_T
	INFO_SPEED_DOWNLOAD       = C.CURLINFO_SPEED_DOWNLOAD
	INFO_SPEED_DOWNLOAD_T     = C.CURLINFO_SPEED_DOWNLOAD_T
	INFO_SPEED_UPLOAD         = C.CURLINFO_SPEED_UPLOAD
	INFO_SPEED_UPLOAD_T       = C.CURLINFO_SPEED_UPLOAD_T
	INFO_HEADER_SIZE          = C.CURLINFO_HEADER_SIZE
	INFO_REQUEST_SIZE         = C.CURLINFO_REQUEST_SIZE
	INFO_SSL_VERIFYRESULT     = C.CURLINFO_SSL_VERIFYRESULT
	INFO_FILETIME             = C.CURLINFO_FILETIME
	INFO_FILETIME_T           = C.CURLINFO_FILETIME_T
	INFO_CONTENT_LENGTH_DOWNLOAD = C.CURLINFO_CONTENT_LENGTH_DOWNLOAD
	INFO_CONTENT_LENGTH_DOWNLOAD_T = C.CURLINFO_CONTENT_LENGTH_DOWNLOAD_T
	INFO_CONTENT_LENGTH_UPLOAD = C.CURLINFO_CONTENT_LENGTH_UPLOAD
	INFO_CONTENT_LENGTH_UPLOAD_T = C.CURLINFO_CONTENT_LENGTH_UPLOAD_T
	INFO_STARTTRANSFER_TIME   = C.CURLINFO_STARTTRANSFER_TIME
	INFO_CONTENT_TYPE         = C.CURLINFO_CONTENT_TYPE
	INFO_REDIRECT_TIME        = C.CURLINFO_REDIRECT_TIME
	INFO_REDIRECT_COUNT       = C.CURLINFO_REDIRECT_COUNT
	INFO_PRIVATE              = C.CURLINFO_PRIVATE
	INFO_HTTP_CONNECTCODE     = C.CURLINFO_HTTP_CONNECTCODE
	INFO_HTTPAUTH_AVAIL       = C.CURLINFO_HTTPAUTH_AVAIL
	INFO_PROXYAUTH_AVAIL      = C.CURLINFO_PROXYAUTH_AVAIL
	INFO_OS_ERRNO             = C.CURLINFO_OS_ERRNO
	INFO_NUM_CONNECTS         = C.CURLINFO_NUM_CONNECTS
	INFO_SSL_ENGINES          = C.CURLINFO_SSL_ENGINES
	INFO_COOKIELIST           = C.CURLINFO_COOKIELIST
	INFO_LASTSOCKET           = C.CURLINFO_LASTSOCKET
	INFO_FTP_ENTRY_PATH       = C.CURLINFO_FTP_ENTRY_PATH
	INFO_REDIRECT_URL         = C.CURLINFO_REDIRECT_URL
	INFO_PRIMARY_IP           = C.CURLINFO_PRIMARY_IP
	INFO_APPCONNECT_TIME      = C.CURLINFO_APPCONNECT_TIME
	INFO_CERTINFO             = C.CURLINFO_CERTINFO
	INFO_CONDITION_UNMET      = C.CURLINFO_CONDITION_UNMET
	INFO_RTSP_SESSION_ID      = C.CURLINFO_RTSP_SESSION_ID
	INFO_RTSP_CLIENT_CSEQ     = C.CURLINFO_RTSP_CLIENT_CSEQ
	INFO_RTSP_SERVER_CSEQ     = C.CURLINFO_RTSP_SERVER_CSEQ
	INFO_RTSP_CSEQ_RECV       = C.CURLINFO_RTSP_CSEQ_RECV
	INFO_PRIMARY_PORT         = C.CURLINFO_PRIMARY_PORT
	INFO_LOCAL_IP             = C.CURLINFO_LOCAL_IP
	INFO_LOCAL_PORT           = C.CURLINFO_LOCAL_PORT
	INFO_TLS_SESSION          = C.CURLINFO_TLS_SESSION
	INFO_ACTIVESOCKET         = C.CURLINFO_ACTIVESOCKET
	INFO_TLS_SSL_PTR          = C.CURLINFO_TLS_SSL_PTR
	INFO_HTTP_VERSION         = C.CURLINFO_HTTP_VERSION
	INFO_PROXY_SSL_VERIFYRESULT = C.CURLINFO_PROXY_SSL_VERIFYRESULT
	INFO_PROTOCOL             = C.CURLINFO_PROTOCOL
	INFO_SCHEME               = C.CURLINFO_SCHEME
	INFO_TOTAL_TIME_T         = C.CURLINFO_TOTAL_TIME_T
	INFO_NAMELOOKUP_TIME_T    = C.CURLINFO_NAMELOOKUP_TIME_T
	INFO_CONNECT_TIME_T       = C.CURLINFO_CONNECT_TIME_T
	INFO_PRETRANSFER_TIME_T   = C.CURLINFO_PRETRANSFER_TIME_T
	INFO_STARTTRANSFER_TIME_T = C.CURLINFO_STARTTRANSFER_TIME_T
	INFO_REDIRECT_TIME_T      = C.CURLINFO_REDIRECT_TIME_T
	INFO_APPCONNECT_TIME_T    = C.CURLINFO_APPCONNECT_TIME_T
	INFO_RETRY_AFTER          = C.CURLINFO_RETRY_AFTER
	INFO_EFFECTIVE_METHOD     = C.CURLINFO_EFFECTIVE_METHOD
	INFO_PROXY_ERROR          = C.CURLINFO_PROXY_ERROR
	INFO_REFERER              = C.CURLINFO_REFERER
	INFO_CAINFO               = C.CURLINFO_CAINFO
	INFO_CAPATH               = C.CURLINFO_CAPATH
	INFO_XFER_ID              = C.CURLINFO_XFER_ID
	INFO_CONN_ID              = C.CURLINFO_CONN_ID
	INFO_QUEUE_TIME_T         = C.CURLINFO_QUEUE_TIME_T
	INFO_USED_PROXY           = C.CURLINFO_USED_PROXY
	INFO_LASTONE              = C.CURLINFO_LASTONE
	INFO_HTTP_CODE            = C.CURLINFO_HTTP_CODE
)

// Auth
const (
	AUTH_NONE                      = C.CURLAUTH_NONE & (1<<32 - 1)
	AUTH_BASIC                     = C.CURLAUTH_BASIC & (1<<32 - 1)
	AUTH_DIGEST                    = C.CURLAUTH_DIGEST & (1<<32 - 1)
	AUTH_NEGOTIATE                 = C.CURLAUTH_NEGOTIATE & (1<<32 - 1)
	AUTH_GSSNEGOTIATE              = C.CURLAUTH_GSSNEGOTIATE & (1<<32 - 1)
	AUTH_GSSAPI                    = C.CURLAUTH_GSSAPI & (1<<32 - 1)
	AUTH_NTLM                      = C.CURLAUTH_NTLM & (1<<32 - 1)
	AUTH_DIGEST_IE                 = C.CURLAUTH_DIGEST_IE & (1<<32 - 1)
	AUTH_NTLM_WB                   = C.CURLAUTH_NTLM_WB & (1<<32 - 1)
	AUTH_BEARER                    = C.CURLAUTH_BEARER & (1<<32 - 1)
	AUTH_AWS_SIGV4                 = C.CURLAUTH_AWS_SIGV4 & (1<<32 - 1)
	AUTH_ONLY                      = C.CURLAUTH_ONLY & (1<<32 - 1)
	AUTH_ANY                       = C.CURLAUTH_ANY & (1<<32 - 1)
	AUTH_ANYSAFE                   = C.CURLAUTH_ANYSAFE & (1<<32 - 1)
)

// generated ends
