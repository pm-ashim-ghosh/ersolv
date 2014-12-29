
============================
 Centralised Error Handling
============================

BootCamp, Dec 2014 assignment for group 5-

- Nikhil Bora
- Ashim Ghosh
- Chirag Rupani
- Sagar Karmarkar

See ``doc/presentation.ppt`` and ``doc/code-insights.rst`` for an overview of
the project.


--------------
 Installation
--------------

Download and install::

  $ go install github.com/pm-ashim-ghosh/ersolv

This should place the code in
``${GOPATH}/src/github.com/pm-ashim-ghosh/ersolv/``.

Run the server::

  $ cd ${GOPATH}/bin && ./ersolv

Visit ``localhost:4000/ersolv`` to fetch a list of "code logs".

Use the following CURL command to add new "code logs" though the API- ::

  $ curl -X POST -H "Content-Type: application/json" -d '{
        "Log_code": "ADSRTB0001",
        "Filepath": "adserver/rtb",
        "Line_no": 10}'
    localhost:4000/ersolv


-------------------
 Future Directions
-------------------

- Implement ``GET /`` to make the API HATEOAS compliant.
- Handle other types of errors: data center errors, business logic, errors,
  more?)
