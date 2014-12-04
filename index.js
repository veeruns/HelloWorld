var http = require('http');

http.createServer(function(request, response) {

    response.setHeader('Connection', 'Transfer-Encoding');
    response.setHeader('Content-Type', 'text/html; charset=utf-8');
    response.setHeader('Transfer-Encoding', 'chunked');

    response.write('hello');

    setTimeout(function() {
        response.write(' world!');
        response.end();
    }, 10000);

}).listen(80);
