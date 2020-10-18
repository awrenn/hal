# hal

hal is a tool made for staticly compiling html snippets into a single html file.  
This tool was made to replace the HtmlWebpackPlugin, which does this, but only for a single layer.

For example:  
Target file ./dist/index.html is based on source file ./src/index.html.  
It has a header bar that is common across all html files. It indicates this by referencing some ./src/header.html.

Using HtmlWebpackPlugin:  
`<%= require("html-loader!./header.html") %>`

Using hal:
`<%header.html %/>`

But if "header.html" wants to include its own html files, the HtmlWebpackPlugin cannot do this.

Using hal, "header.html" could include the line:
`<%logo.html %/>`

Which would then inject "./src/logo.html".

Obviously, the surrounding feature set is lacking compared to HtmlWebpackPlugin, but this tool was made to solve the author's problems first.

## Basic Usage  

Currently, usage is fairly static. Configuration is the next goal.    
The user provides a yaml file that indicates what the target files are. The user also provides a src directory, and a dst directory.  
Example:  
`hal ./input.yaml --src ./src --dst ./dst`  
"input.yaml" provides a list of files, which are then streamed from ./src/\<file\>, and saved to ./dst/\<file\>.  

Currently, hal attacks js-beautify in html mode onto the output stream, instead of handling pretty-printing itself.

## How it works

hal aims to be as memory efficient as possible first, and then fast second, and then feature-ful third. It works by opening the input and output files as streams, and then parsining each tag at a time. Each time a reference tag is parsed, a tag that takes the form `<%file.html%/>`, hal pauses the stream, and starts streaming the referenced file instead. hal keeps a history of all referenced files, which allows it to detect cycles, and throw an error.  
This means that memory usage of hal itself depends on the number of files, for the history component, and on the size of the largest tag, since it keeps each tag in memory + an extra byte buffer as it reads from each source stream.


