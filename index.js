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
       var data=listChildren(client, path);
       var i=0;
       reponse.write("data</br>");
    });
    response.end();	

}).listen(80);



function listChildren(client, path) {
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
            console.log('Children of %s are: %j.', path, children);
            return children;
		
        }
    );
}


