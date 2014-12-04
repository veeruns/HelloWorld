var fs=require('fs');
var name_of_file=process.argv[2];
var buffer=fs.readFileSync(name_of_file);
var text=buffer.toString();
var splitted=text.split("\n");
var num_new_line=splitted.length;
num_new_line=num_new_line - 1;
console.log(num_new_line);



