var fs=require('fs');
var name_of_file=process.argv[2];
var buffer=fs.readFile(name_of_file,callback);



function callback(err,data){
if(err) throw err;
var text=data.toString();
var splitted=text.split("\n");
var num_new_line=splitted.length;
num_new_line=num_new_line - 1;
console.log(num_new_line);
}


