(()=>{if("undefined"==typeof global)if("undefined"!=typeof window)window.global=window;else{if("undefined"==typeof self)throw new Error("cannot export Go (neither global, window nor self is defined)");self.global=self}global.require||"undefined"==typeof require||(global.require=require),!global.fs&&global.require&&(global.fs=require("fs"));const enosys=()=>{const err=new Error("not implemented");return err.code="ENOSYS",err};if(!global.fs){let outputBuf="";global.fs={constants:{O_WRONLY:-1,O_RDWR:-1,O_CREAT:-1,O_TRUNC:-1,O_APPEND:-1,O_EXCL:-1},writeSync(fd,buf){var nl=(outputBuf+=decoder.decode(buf)).lastIndexOf("\n");return-1!=nl&&(outputBuf=outputBuf.substr(nl+1)),buf.length},write(fd,buf,offset,length,position,callback){0!==offset||length!==buf.length||null!==position?callback(enosys()):callback(null,this.writeSync(fd,buf))},chmod(path,mode,callback){callback(enosys())},chown(path,uid,gid,callback){callback(enosys())},close(fd,callback){callback(enosys())},fchmod(fd,mode,callback){callback(enosys())},fchown(fd,uid,gid,callback){callback(enosys())},fstat(fd,callback){callback(enosys())},fsync(fd,callback){callback(null)},ftruncate(fd,length,callback){callback(enosys())},lchown(path,uid,gid,callback){callback(enosys())},link(path,link,callback){callback(enosys())},lstat(path,callback){callback(enosys())},mkdir(path,perm,callback){callback(enosys())},open(path,flags,mode,callback){callback(enosys())},read(fd,buffer,offset,length,position,callback){callback(enosys())},readdir(path,callback){callback(enosys())},readlink(path,callback){callback(enosys())},rename(from,to,callback){callback(enosys())},rmdir(path,callback){callback(enosys())},stat(path,callback){callback(enosys())},symlink(path,link,callback){callback(enosys())},truncate(path,length,callback){callback(enosys())},unlink(path,callback){callback(enosys())},utimes(path,atime,mtime,callback){callback(enosys())}}}if(global.process||(global.process={getuid(){return-1},getgid(){return-1},geteuid(){return-1},getegid(){return-1},getgroups(){throw enosys()},pid:-1,ppid:-1,umask(){throw enosys()},cwd(){throw enosys()},chdir(){throw enosys()}}),!global.crypto){const nodeCrypto=require("crypto");global.crypto={getRandomValues(b){nodeCrypto.randomFillSync(b)}}}global.performance||(global.performance={now(){var[sec,nsec]=process.hrtime();return 1e3*sec+nsec/1e6}}),global.TextEncoder||(global.TextEncoder=require("util").TextEncoder),global.TextDecoder||(global.TextDecoder=require("util").TextDecoder);const encoder=new TextEncoder("utf-8"),decoder=new TextDecoder("utf-8");var logLine=[];if(global.Go=class{constructor(){this._callbackTimeouts=new Map,this._nextCallbackTimeoutID=1;const mem=()=>new DataView(this._inst.exports.memory.buffer),setInt64=(addr,v)=>{mem().setUint32(addr+0,v,!0),mem().setUint32(addr+4,Math.floor(v/4294967296),!0)};const loadValue=addr=>{var f=mem().getFloat64(addr,!0);if(0!==f)return isNaN(f)?(addr=mem().getUint32(addr,!0),this._values[addr]):f},storeValue=(addr,v)=>{if("number"==typeof v)return isNaN(v)?(mem().setUint32(addr+4,2146959360,!0),void mem().setUint32(addr,0,!0)):0===v?(mem().setUint32(addr+4,2146959360,!0),void mem().setUint32(addr,1,!0)):void mem().setFloat64(addr,v,!0);switch(v){case void 0:return void mem().setFloat64(addr,0,!0);case null:return mem().setUint32(addr+4,2146959360,!0),void mem().setUint32(addr,2,!0);case!0:return mem().setUint32(addr+4,2146959360,!0),void mem().setUint32(addr,3,!0);case!1:return mem().setUint32(addr+4,2146959360,!0),void mem().setUint32(addr,4,!0)}let id=this._ids.get(v),typeFlag=(void 0===id&&(void 0===(id=this._idPool.pop())&&(id=this._values.length),this._values[id]=v,this._goRefCounts[id]=0,this._ids.set(v,id)),this._goRefCounts[id]++,1);switch(typeof v){case"string":typeFlag=2;break;case"symbol":typeFlag=3;break;case"function":typeFlag=4}mem().setUint32(addr+4,2146959360|typeFlag,!0),mem().setUint32(addr,id,!0)},loadSlice=(array,len,cap)=>new Uint8Array(this._inst.exports.memory.buffer,array,len),loadSliceOfValues=(array,len,cap)=>{const a=new Array(len);for(let i=0;i<len;i++)a[i]=loadValue(array+8*i);return a},loadString=(ptr,len)=>decoder.decode(new DataView(this._inst.exports.memory.buffer,ptr,len)),timeOrigin=Date.now()-performance.now();this.importObject={wasi_snapshot_preview1:{fd_write:function(fd,iovs_ptr,iovs_len,nwritten_ptr){let nwritten=0;if(1==fd)for(let iovs_i=0;iovs_i<iovs_len;iovs_i++){var iov_ptr=iovs_ptr+8*iovs_i,ptr=mem().getUint32(iov_ptr+0,!0),len=mem().getUint32(iov_ptr+4,!0);nwritten+=len;for(let i=0;i<len;i++){var c=mem().getUint8(ptr+i);13!=c&&(10==c?(decoder.decode(new Uint8Array(logLine)),logLine=[]):logLine.push(c))}}return mem().setUint32(nwritten_ptr,nwritten,!0),0},fd_close:()=>0,fd_fdstat_get:()=>0,fd_seek:()=>0,proc_exit:code=>{if(!global.process)throw"trying to exit with code "+code;process.exit(code)},random_get:(bufPtr,bufLen)=>(crypto.getRandomValues(loadSlice(bufPtr,bufLen)),0)},env:{"runtime.ticks":()=>timeOrigin+performance.now(),"runtime.sleepTicks":timeout=>{setTimeout(this._inst.exports.go_scheduler,timeout)},"syscall/js.finalizeRef":sp=>{},"syscall/js.stringVal":(ret_ptr,value_ptr,value_len)=>{value_ptr=loadString(value_ptr,value_len);storeValue(ret_ptr,value_ptr)},"syscall/js.valueGet":(retval,v_addr,p_ptr,p_len)=>{p_ptr=loadString(p_ptr,p_len),p_len=loadValue(v_addr),v_addr=Reflect.get(p_len,p_ptr);storeValue(retval,v_addr)},"syscall/js.valueSet":(v_addr,p_ptr,p_len,x_addr)=>{v_addr=loadValue(v_addr),p_ptr=loadString(p_ptr,p_len),p_len=loadValue(x_addr);Reflect.set(v_addr,p_ptr,p_len)},"syscall/js.valueDelete":(v_addr,p_ptr,p_len)=>{v_addr=loadValue(v_addr),p_ptr=loadString(p_ptr,p_len);Reflect.deleteProperty(v_addr,p_ptr)},"syscall/js.valueIndex":(ret_addr,v_addr,i)=>{storeValue(ret_addr,Reflect.get(loadValue(v_addr),i))},"syscall/js.valueSetIndex":(v_addr,i,x_addr)=>{Reflect.set(loadValue(v_addr),i,loadValue(x_addr))},"syscall/js.valueCall":(ret_addr,v_addr,m_ptr,m_len,args_ptr,args_len,args_cap)=>{v_addr=loadValue(v_addr),m_ptr=loadString(m_ptr,m_len),m_len=loadSliceOfValues(args_ptr,args_len);try{var m=Reflect.get(v_addr,m_ptr);storeValue(ret_addr,Reflect.apply(m,v_addr,m_len)),mem().setUint8(ret_addr+8,1)}catch(err){storeValue(ret_addr,err),mem().setUint8(ret_addr+8,0)}},"syscall/js.valueInvoke":(ret_addr,v_addr,args_ptr,args_len,args_cap)=>{try{var v=loadValue(v_addr),args=loadSliceOfValues(args_ptr,args_len);storeValue(ret_addr,Reflect.apply(v,void 0,args)),mem().setUint8(ret_addr+8,1)}catch(err){storeValue(ret_addr,err),mem().setUint8(ret_addr+8,0)}},"syscall/js.valueNew":(ret_addr,v_addr,args_ptr,args_len,args_cap)=>{v_addr=loadValue(v_addr),args_ptr=loadSliceOfValues(args_ptr,args_len);try{storeValue(ret_addr,Reflect.construct(v_addr,args_ptr)),mem().setUint8(ret_addr+8,1)}catch(err){storeValue(ret_addr,err),mem().setUint8(ret_addr+8,0)}},"syscall/js.valueLength":v_addr=>loadValue(v_addr).length,"syscall/js.valuePrepareString":(ret_addr,v_addr)=>{v_addr=String(loadValue(v_addr)),v_addr=encoder.encode(v_addr);storeValue(ret_addr,v_addr),setInt64(ret_addr+8,v_addr.length)},"syscall/js.valueLoadString":(v_addr,slice_ptr,slice_len,slice_cap)=>{v_addr=loadValue(v_addr);loadSlice(slice_ptr,slice_len).set(v_addr)},"syscall/js.valueInstanceOf":(v_addr,t_addr)=>loadValue(v_addr)instanceof loadValue(t_addr),"syscall/js.copyBytesToGo":(ret_addr,dest_addr,dest_len,dest_cap,source_addr)=>{var num_bytes_copied_addr=ret_addr,ret_addr=ret_addr+4;const dst=loadSlice(dest_addr,dest_len),src=loadValue(source_addr);src instanceof Uint8Array||src instanceof Uint8ClampedArray?(dest_addr=src.subarray(0,dst.length),dst.set(dest_addr),setInt64(num_bytes_copied_addr,dest_addr.length),mem().setUint8(ret_addr,1)):mem().setUint8(ret_addr,0)},"syscall/js.copyBytesToJS":(ret_addr,dest_addr,source_addr,source_len,source_cap)=>{var num_bytes_copied_addr=ret_addr,ret_addr=ret_addr+4;const dst=loadValue(dest_addr),src=loadSlice(source_addr,source_len);dst instanceof Uint8Array||dst instanceof Uint8ClampedArray?(dest_addr=src.subarray(0,dst.length),dst.set(dest_addr),setInt64(num_bytes_copied_addr,dest_addr.length),mem().setUint8(ret_addr,1)):mem().setUint8(ret_addr,0)}}}}async run(instance){this._inst=instance,this._values=[NaN,0,null,!0,!1,global,this],this._goRefCounts=[],this._ids=new Map,this._idPool=[],this.exited=!1;for(new DataView(this._inst.exports.memory.buffer);;){var callbackPromise=new Promise(resolve=>{this._resolveCallbackPromise=()=>{if(this.exited)throw new Error("bad callback: Go program has already exited");setTimeout(resolve,0)}});if(this._inst.exports._start(),this.exited)break;await callbackPromise}}_resume(){if(this.exited)throw new Error("Go program has already exited");this._inst.exports.resume(),this.exited&&this._resolveExitPromise()}_makeFuncWrapper(id){const go=this;return function(){var event={id:id,this:this,args:arguments};return go._pendingEvent=event,go._resume(),event.result}}},global.require&&global.require.main===module&&global.process&&global.process.versions&&!global.process.versions.electron){3!=process.argv.length&&process.exit(1);const go=new Go;WebAssembly.instantiate(fs.readFileSync(process.argv[2]),go.importObject).then(result=>go.run(result.instance)).catch(err=>{process.exit(1)})}})();