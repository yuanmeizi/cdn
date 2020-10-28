# 

># Hypertext Transfer Protocol (HTTP) Status Code Registry

# htttp 状态码库

 参照:https://www.iana.org/assignments/http-status-codes/http-status-codes.xml

   //HttpStatus //ResponseEntity

- Last Updated

  2018-09-21

- Available Formats

  [![img](https://www.iana.org/_img/icons/text-xml.png) XML](https://www.iana.org/assignments/http-status-codes/http-status-codes.xml)[![img](https://www.iana.org/_img/icons/text-html.png) HTML](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml)[![img](https://www.iana.org/_img/icons/text-plain.png) Plain text](https://www.iana.org/assignments/http-status-codes/http-status-codes.txt)

**Registry included below**

- [HTTP Status Codes](https://www.iana.org/assignments/http-status-codes/http-status-codes.xml#http-status-codes-1)

HTTP Status Codes

- Registration Procedure(s)

  `IETF Review`

- Reference

  [[RFC7231](https://www.iana.org/go/rfc7231)]

- Note

  `1xx: Informational - Request received, continuing process 2xx: Success - The action was successfully received, understood, and accepted 3xx: Redirection - Further action must be taken in order to complete the request 4xx: Client Error - The request contains bad syntax or cannot be fulfilled 5xx: Server Error - The server failed to fulfill an apparently valid request    `

- Available Formats

  [![img](https://www.iana.org/_img/icons/text-csv.png) CSV](https://www.iana.org/assignments/http-status-codes/http-status-codes-1.csv)

| Value ![img](https://www.iana.org/assignments/_support/sort_none.gif) | Description ![img](https://www.iana.org/assignments/_support/sort_none.gif) | Reference ![img](https://www.iana.org/assignments/_support/sort_none.gif) |
| :----------------------------------------------------------- | :----------------------------------------------------------- | :----------------------------------------------------------- |
| 100                                                          | Continue                                                     | [[RFC7231, Section 6.2.1](https://www.iana.org/go/rfc7231)]  |
| 101                                                          | Switching Protocols                                          | [[RFC7231, Section 6.2.2](https://www.iana.org/go/rfc7231)]  |
| 102                                                          | Processing                                                   | [[RFC2518](https://www.iana.org/go/rfc2518)]                 |
| 103                                                          | Early Hints                                                  | [[RFC8297](https://www.iana.org/go/rfc8297)]                 |
| 104-199                                                      | Unassigned                                                   |                                                              |
| 200                                                          | OK                                                           | [[RFC7231, Section 6.3.1](https://www.iana.org/go/rfc7231)]  |
| 201                                                          | Created                                                      | [[RFC7231, Section 6.3.2](https://www.iana.org/go/rfc7231)]  |
| 202                                                          | Accepted                                                     | [[RFC7231, Section 6.3.3](https://www.iana.org/go/rfc7231)]  |
| 203                                                          | Non-Authoritative Information                                | [[RFC7231, Section 6.3.4](https://www.iana.org/go/rfc7231)]  |
| 204                                                          | No Content                                                   | [[RFC7231, Section 6.3.5](https://www.iana.org/go/rfc7231)]  |
| 205                                                          | Reset Content                                                | [[RFC7231, Section 6.3.6](https://www.iana.org/go/rfc7231)]  |
| 206                                                          | Partial Content                                              | [[RFC7233, Section 4.1](https://www.iana.org/go/rfc7233)]    |
| 207                                                          | Multi-Status                                                 | [[RFC4918](https://www.iana.org/go/rfc4918)]                 |
| 208                                                          | Already Reported                                             | [[RFC5842](https://www.iana.org/go/rfc5842)]                 |
| 209-225                                                      | Unassigned                                                   |                                                              |
| 226                                                          | IM Used                                                      | [[RFC3229](https://www.iana.org/go/rfc3229)]                 |
| 227-299                                                      | Unassigned                                                   |                                                              |
| 300                                                          | Multiple Choices                                             | [[RFC7231, Section 6.4.1](https://www.iana.org/go/rfc7231)]  |
| 301                                                          | Moved Permanently                                            | [[RFC7231, Section 6.4.2](https://www.iana.org/go/rfc7231)]  |
| 302                                                          | Found                                                        | [[RFC7231, Section 6.4.3](https://www.iana.org/go/rfc7231)]  |
| 303                                                          | See Other                                                    | [[RFC7231, Section 6.4.4](https://www.iana.org/go/rfc7231)]  |
| 304                                                          | Not Modified                                                 | [[RFC7232, Section 4.1](https://www.iana.org/go/rfc7232)]    |
| 305                                                          | Use Proxy                                                    | [[RFC7231, Section 6.4.5](https://www.iana.org/go/rfc7231)]  |
| 306                                                          | (Unused)                                                     | [[RFC7231, Section 6.4.6](https://www.iana.org/go/rfc7231)]  |
| 307                                                          | Temporary Redirect                                           | [[RFC7231, Section 6.4.7](https://www.iana.org/go/rfc7231)]  |
| 308                                                          | Permanent Redirect                                           | [[RFC7538](https://www.iana.org/go/rfc7538)]                 |
| 309-399                                                      | Unassigned                                                   |                                                              |
| 400                                                          | Bad Request                                                  | [[RFC7231, Section 6.5.1](https://www.iana.org/go/rfc7231)]  |
| 401                                                          | Unauthorized                                                 | [[RFC7235, Section 3.1](https://www.iana.org/go/rfc7235)]    |
| 402                                                          | Payment Required                                             | [[RFC7231, Section 6.5.2](https://www.iana.org/go/rfc7231)]  |
| 403                                                          | Forbidden                                                    | [[RFC7231, Section 6.5.3](https://www.iana.org/go/rfc7231)]  |
| 404                                                          | Not Found                                                    | [[RFC7231, Section 6.5.4](https://www.iana.org/go/rfc7231)]  |
| 405                                                          | Method Not Allowed                                           | [[RFC7231, Section 6.5.5](https://www.iana.org/go/rfc7231)]  |
| 406                                                          | Not Acceptable                                               | [[RFC7231, Section 6.5.6](https://www.iana.org/go/rfc7231)]  |
| 407                                                          | Proxy Authentication Required                                | [[RFC7235, Section 3.2](https://www.iana.org/go/rfc7235)]    |
| 408                                                          | Request Timeout                                              | [[RFC7231, Section 6.5.7](https://www.iana.org/go/rfc7231)]  |
| 409                                                          | Conflict                                                     | [[RFC7231, Section 6.5.8](https://www.iana.org/go/rfc7231)]  |
| 410                                                          | Gone                                                         | [[RFC7231, Section 6.5.9](https://www.iana.org/go/rfc7231)]  |
| 411                                                          | Length Required                                              | [[RFC7231, Section 6.5.10](https://www.iana.org/go/rfc7231)] |
| 412                                                          | Precondition Failed                                          | [[RFC7232, Section 4.2](https://www.iana.org/go/rfc7232)][[RFC8144, Section 3.2](https://www.iana.org/go/rfc8144)] |
| 413                                                          | Payload Too Large                                            | [[RFC7231, Section 6.5.11](https://www.iana.org/go/rfc7231)] |
| 414                                                          | URI Too Long                                                 | [[RFC7231, Section 6.5.12](https://www.iana.org/go/rfc7231)] |
| 415                                                          | Unsupported Media Type                                       | [[RFC7231, Section 6.5.13](https://www.iana.org/go/rfc7231)][[RFC7694, Section 3](https://www.iana.org/go/rfc7694)] |
| 416                                                          | Range Not Satisfiable                                        | [[RFC7233, Section 4.4](https://www.iana.org/go/rfc7233)]    |
| 417                                                          | Expectation Failed                                           | [[RFC7231, Section 6.5.14](https://www.iana.org/go/rfc7231)] |
| 418-420                                                      | Unassigned                                                   |                                                              |
| 421                                                          | Misdirected Request                                          | [[RFC7540, Section 9.1.2](https://www.iana.org/go/rfc7540)]  |
| 422                                                          | Unprocessable Entity                                         | [[RFC4918](https://www.iana.org/go/rfc4918)]                 |
| 423                                                          | Locked                                                       | [[RFC4918](https://www.iana.org/go/rfc4918)]                 |
| 424                                                          | Failed Dependency                                            | [[RFC4918](https://www.iana.org/go/rfc4918)]                 |
| 425                                                          | Too Early                                                    | [[RFC8470](https://www.iana.org/go/rfc8470)]                 |
| 426                                                          | Upgrade Required                                             | [[RFC7231, Section 6.5.15](https://www.iana.org/go/rfc7231)] |
| 427                                                          | Unassigned                                                   |                                                              |
| 428                                                          | Precondition Required                                        | [[RFC6585](https://www.iana.org/go/rfc6585)]                 |
| 429                                                          | Too Many Requests                                            | [[RFC6585](https://www.iana.org/go/rfc6585)]                 |
| 430                                                          | Unassigned                                                   |                                                              |
| 431                                                          | Request Header Fields Too Large                              | [[RFC6585](https://www.iana.org/go/rfc6585)]                 |
| 432-450                                                      | Unassigned                                                   |                                                              |
| 451                                                          | Unavailable For Legal Reasons                                | [[RFC7725](https://www.iana.org/go/rfc7725)]                 |
| 452-499                                                      | Unassigned                                                   |                                                              |
| 500                                                          | Internal Server Error                                        | [[RFC7231, Section 6.6.1](https://www.iana.org/go/rfc7231)]  |
| 501                                                          | Not Implemented                                              | [[RFC7231, Section 6.6.2](https://www.iana.org/go/rfc7231)]  |
| 502                                                          | Bad Gateway                                                  | [[RFC7231, Section 6.6.3](https://www.iana.org/go/rfc7231)]  |
| 503                                                          | Service Unavailable                                          | [[RFC7231, Section 6.6.4](https://www.iana.org/go/rfc7231)]  |
| 504                                                          | Gateway Timeout                                              | [[RFC7231, Section 6.6.5](https://www.iana.org/go/rfc7231)]  |
| 505                                                          | HTTP Version Not Supported                                   | [[RFC7231, Section 6.6.6](https://www.iana.org/go/rfc7231)]  |
| 506                                                          | Variant Also Negotiates                                      | [[RFC2295](https://www.iana.org/go/rfc2295)]                 |
| 507                                                          | Insufficient Storage                                         | [[RFC4918](https://www.iana.org/go/rfc4918)]                 |
| 508                                                          | Loop Detected                                                | [[RFC5842](https://www.iana.org/go/rfc5842)]                 |
| 509                                                          | Unassigned                                                   |                                                              |
| 510                                                          | Not Extended                                                 | [[RFC2774](https://www.iana.org/go/rfc2774)]                 |
| 511                                                          | Network Authentication Required                              | [[RFC6585](https://www.iana.org/go/rfc6585)]                 |
| 512-599                                                      | Unassigned                                                   |                                                              |