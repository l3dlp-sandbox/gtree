const go=new Go;let mod,instance;WebAssembly.instantiateStreaming(fetch("main.wasm"),go.importObject).then(result=>{mod=result.module,instance=result.instance,document.getElementById("gtree").disabled=!1,go.run(instance),instance=WebAssembly.instantiate(mod,go.importObject)});const clearMarkdown=()=>{document.getElementById("in").value=""},generateTree=()=>{gtree()},copyToClipboard=()=>{var tree=document.getElementById("treeView");if(null!==tree){const clipboard=window.navigator.clipboard;clipboard.writeText(tree.innerHTML)}};