var http = require('http');
var zookeeper = require('node-zookeeper-client');
var client = zookeeper.createClient('zk101.iris.ne1.yahoo.com:2181');
var path = '/prod/services';

http.createServer(function(request, response) {

    response.setHeader('Connection', 'Transfer-Encoding');
    response.setHeader('Content-Type', 'text/html; charset=utf-8');
    client.connect();
    client.once('connected', function () {
       response.write('Connected to ZooKeeper.');
       listChildren(client, path,response);
    });
    response.end();	

}).listen(80);



function listChildren(client, path,response) {
    client.getChildren(
        path,
        function (error, children, stat) {
            if (error) {
                  console.log(
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


