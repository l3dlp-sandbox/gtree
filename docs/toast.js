const toast=document.getElementById("toast");let isVisible=!1;const showToast=(status,message)=>{if(isVisible)return!1;toast.innerHTML=message,toast.classList.add("is-"+status),toast.classList.add("is-show"),isVisible=!0};toast.addEventListener("animationend",()=>{toast.innerHTML="",toast.className="init",isVisible=!1}),document.getElementById("copy").addEventListener("click",ev=>{showToast("success","Copied!")});