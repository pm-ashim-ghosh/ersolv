
------------------------
 Evolving Understanding
------------------------

-----------------------
 Alternative Solutions
-----------------------

- Redirections

  - Server Load
  - Rewrite all errors
  - Complex error formats (different languages)
  - No error resolution

- RuleEngine

  - Error Resolution

- Ersolv: Client-Server over HTTP REST API with RuleEngine

  - Scalable (seperation of concerns)
  - JSON error formats


--------------
 Architecture
--------------

TODO


---------------
 HTTP REST API
---------------

Sources of errors-

- Code Runtime errors
- Data Center errors
- Business Logic errors
- More ???

Resource Model
==============

Logging logs/errors- ::

  POST /ersolv
  {
      "Log_code":"ADSRTB0001",
      "Filepath":"ad_server_main.c",
      "Line_no":10,

      // en_us
      "Message":"Ersolve Test:: Error at 10 in ad_server_main.c"
  }

Server response::

  {
    id: 1
  }

View logs and their status- ::

  GET /ersolv
  {
      "Log_id":1,
      "Log_code":"ADSRTB0001",
      "Type":"ersolv-test",
      "Severity":2,
      "Source":"adserver/rtb",
      "Filepath":"ad_server_main.c",
      "Line_no":10,
      "Message":"Ersolve Test:: Error at 10 in ad_server_main.c"
  }


An attempt at HATEOAS compatiability- ::

  GET /

  {
      "version": 0.1

      "create": {
          "supports": true
          "form": {
              "method": "POST"
              "action": "/ersolve"
              "type": "log"
              "fields": [
                  {
                      "name": "log_code",
                      "type": "string"
                  },
                  {
                      "name": "filepath",
                      "type": "string"
                  },
                  {
                      "name": "line_no",
                      "type": "int"
                  },
              ]
          }
      }

      "read": {
          "supports": true
          "form": {
              "method": "GET"
              "action": "/ersolve"
              "type": "log"
              "fields": [
                  {
                      "name": "log_id",
                      "type": "int"
                  },
                  {
                      "name": "log_code",
                      "type": "string"
                  },
                  {
                      "name": "type",
                      "type": "string"
                  },
                  {
                      "name": "severity",
                      "type": "int"
                  },
                  {
                      "name": "source",
                      "type": "string"
                  },
                  {
                      "name": "filepath",
                      "type": "string"
                  },
                  {
                      "name": "line_no",
                      "type": "int"
                  },
                  {
                      "name": "message",
                      "type": "string"
                  },
              ]
          }
      }
  }

Architectural Constraints
=========================

- Client–server
- Stateless
- Cacheable
- Layered system
- Code on demand (optional)
- Uniform interface

  - Identification of resources
  - Manipulation of resources through these representations
  - Self-descriptive messages
  - Hypermedia as the engine of application state (HATEOAS)


------------
 Data Model
------------

::

  mysql> show tables;
  +------------------+
  | Tables_in_ersolv |
  +------------------+
  | code_logs        |
  | locale           |
  | log_types        |
  | logs             |
  | messages         |
  +------------------+

  mysql> describe logs;
  +----------+------------+------+-----+---------+----------------+
  | Field    | Type       | Null | Key | Default | Extra          |
  +----------+------------+------+-----+---------+----------------+
  | log_id   | bigint(20) | NO   | PRI | NULL    | auto_increment |
  | log_code | char(10)   | YES  | MUL | NULL    |                |
  +----------+------------+------+-----+---------+----------------+

  mysql> describe log_types;
  +----------+-----------+------+-----+---------+-------+
  | Field    | Type      | Null | Key | Default | Extra |
  +----------+-----------+------+-----+---------+-------+
  | log_code | char(10)  | NO   | PRI | NULL    |       |
  | type     | char(20)  | NO   |     | NULL    |       |
  | severity | int(11)   | YES  |     | NULL    |       |
  | source   | char(100) | YES  |     | NULL    |       |
  +----------+-----------+------+-----+---------+-------+

  mysql> describe code_logs;
  +------------+------------+------+-----+---------+-------+
  | Field      | Type       | Null | Key | Default | Extra |
  +------------+------------+------+-----+---------+-------+
  | log_id     | bigint(20) | NO   | PRI | NULL    |       |
  | filepath   | text       | YES  |     | NULL    |       |
  | line_no    | int(11)    | YES  |     | NULL    |       |
  | message_id | int(11)    | YES  | MUL | NULL    |       |
  +------------+------------+------+-----+---------+-------+

  mysql> describe locale;
  +--------------+----------+------+-----+---------+-------+
  | Field        | Type     | Null | Key | Default | Extra |
  +--------------+----------+------+-----+---------+-------+
  | locale_id    | char(10) | NO   | PRI | NULL    |       |
  | currency     | char(20) | YES  |     | NULL    |       |
  | time_format  | char(20) | YES  |     | NULL    |       |
  | number_units | int(11)  | YES  |     | NULL    |       |
  +--------------+----------+------+-----+---------+-------+

  mysql> describe messages;
  +----------------+---------------+------+-----+---------+-------+
  | Field          | Type          | Null | Key | Default | Extra |
  +----------------+---------------+------+-----+---------+-------+
  | message_id     | int(11)       | NO   | PRI | 0       |       |
  | locale_id      | char(10)      | NO   | PRI |         |       |
  | message_string | varchar(1000) | YES  |     | NULL    |       |
  +----------------+---------------+------+-----+---------+-------+


-------------------
 Future Directions
-------------------

- ???
