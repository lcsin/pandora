import{c,a,j as l,v as n,u as d,i as p,g,k as w}from"./index-D32rqE15.js";import{a as _,t as u}from"./index-caoARxxB.js";import{_ as v}from"./_plugin-vue_export-helper-DlAUqK2U.js";const R={class:"registration-container"},y={__name:"Register",setup(k){const m=w();var o,r,t;function f(){_.users.register(o,r,t).then(i=>{i.code==0?(u.NotifySuccess("注册成功"),m.push("/login")):u.NotifyError(i.message)})}return(i,e)=>(g(),c("div",R,[e[3]||(e[3]=a("h2",null,"Register",-1)),l(a("input",{id:"email",type:"email",placeholder:"Email","onUpdate:modelValue":e[0]||(e[0]=s=>p(o)?o.value=s:o=s)},null,512),[[n,d(o)]]),l(a("input",{id:"password",type:"password",placeholder:"Password","onUpdate:modelValue":e[1]||(e[1]=s=>p(r)?r.value=s:r=s)},null,512),[[n,d(r)]]),l(a("input",{id:"confirmPassword",type:"password",placeholder:"Confirm Password","onUpdate:modelValue":e[2]||(e[2]=s=>p(t)?t.value=s:t=s)},null,512),[[n,d(t)]]),a("button",{onClick:f},"Register")]))}},b=v(y,[["__scopeId","data-v-b8bf6022"]]);export{b as default};
