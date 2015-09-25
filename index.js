var http = require('http');
var zookeeper = require('node-zookeeper-client');
var client = zookeeper.createClient('zk101.iris.ne1.yahoo.com:2181');
var path = '/prod/services';

http.createServer(function(request, response) {

    response.setHeader('Connection', 'Transfer-Encoding');
    response.setHeader('Content-Type', 'text/html; charset=utf-8');
    client.connect();
    response.write('hello world\n');
    client.once('connected', function () {
       response.write('Connected to ZooKeeper.');
       listChildren(client, path);
    });
	

}).listen(80);



function listChildren(client, path) {
    client.getChildren(
        path,
        function (event) {
            response.write('Got watcher event: %s', event);
            listChildren(client, path);
        },
        function (error, children, stat) {
            if (error) {
                  response.write(
                    'Failed to list children of %s due to: %s.',
                    path,
                    error
                );
                return;
            }

            response.write('Children of %s are: %j.', path, children);
        }
    );
}


