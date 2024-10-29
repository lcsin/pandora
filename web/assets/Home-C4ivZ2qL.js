import{o as I,r as n,c as P,a as l,b as o,u as E,i as tt,w as lt,F as S,d as et,e as a,f as y,g as D,t as $,h as L}from"./index-D32rqE15.js";import{a as T,c as R,t as _}from"./index-caoARxxB.js";import{_ as nt}from"./_plugin-vue_export-helper-DlAUqK2U.js";const ot={class:"main-content"},at={class:"search-bar"},st={class:"song-list"},ut=["onDblclick"],it={class:"player"},rt={class:"playing-info"},dt={class:"now-playing"},ft={class:"player-controls"},ct={class:"dialog-footer"},pt={__name:"Home",setup(vt){I(()=>{b()});let k=n([]),f=n([]);function b(){T.music.getMusicList().then(e=>{e.code==0?(k.value=e.data,f.value=e.data.map(t=>({title:t.name+" - "+t.author,url:R.baseURL+t.url}))):_.NotifyError(e.message)})}let p=n("");function F(){T.music.searchMusic(p.value).then(e=>{e.code==0?(k.value=e.data,f.value=e.data.map(t=>({title:t.name+" - "+t.author,url:R.baseURL+t.url}))):_.NotifyError(e.message)})}const v=n(!1),H=n(null),m=n([]);function V(){v.value=!1,m.value=[]}function K(){T.music.uploadMusic(m.value).then(e=>{e.code==0?(_.NotifySuccess("上传成功"),v.value=!1,b()):(alert("上传失败"),_.NotifyError(e.message))})}const i=n(null),q=n(null),g=n(null),W=n(null),X=n(null),C=n(null),N=n(null),h=n(null),w=n(null);let r=n(0),U=n(!1);function c(e){e>=0&&e<f.value.length&&(r=e,i.value.src=f.value[r].url,w.value.textContent=f.value[r].title,i.value.play(),g.value.textContent="⏸️")}function j(){i.value.paused?(i.value.play(),g.value.textContent="⏸️"):(i.value.pause(),g.value.textContent="▶️")}function z(){U.value=!0,r=Math.floor(Math.random()*f.value.length),c(r)}function A(){c(r-1)}function G(){c(r+1)}function J(e){const t=(e.pageX-C.value.offsetLeft)/C.value.offsetWidth;i.value.currentTime=t*i.value.duration}function O(){const e=i.value.currentTime,t=i.value.duration,M=e/t*100;N.value.style.width=`${M}%`;const u=d=>{const x=Math.floor(d/60),B=Math.floor(d%60).toString().padStart(2,"0");return`${x}:${B}`};h.value.textContent=`${u(e)} / ${u(t)}`}function Q(){U?(r=Math.floor(Math.random()*f.value.length),c(r)):c(r+1)}function Y(e){c(e)}return(e,t)=>{const M=y("el-input"),u=y("el-button"),d=y("el-tooltip"),x=y("el-upload"),B=y("el-dialog");return D(),P(S,null,[l("div",{class:"sidebar",id:"sidebar"},[t[5]||(t[5]=l("h1",null,"Pandora",-1)),l("ul",null,[l("li",{onClick:b},"💖 我的音乐")])]),l("div",{class:"divider",id:"divider",onMousedown:t[0]||(t[0]=(...s)=>e.clickDrag&&e.clickDrag(...s))},null,32),l("div",ot,[l("div",at,[o(M,{modelValue:E(p),"onUpdate:modelValue":t[1]||(t[1]=s=>tt(p)?p.value=s:p=s),placeholder:"搜索音乐...",clearable:"",onKeyup:lt(F,["enter"])},null,8,["modelValue"])]),l("div",st,[l("table",null,[l("tbody",null,[t[6]||(t[6]=l("tr",null,[l("th",null,"歌曲"),l("th",null,"歌手")],-1)),(D(!0),P(S,null,et(E(k),(s,Z)=>(D(),P("tr",{onDblclick:yt=>Y(Z)},[l("td",null,$(s.name),1),l("td",null,$(s.author),1)],40,ut))),256))])])]),l("div",it,[l("div",{class:"progress-bar",id:"progress-bar",ref_key:"progressBar",ref:C,onClick:J},[l("div",{class:"progress",id:"progress",ref_key:"progress",ref:N},null,512)],512),l("div",rt,[l("div",dt,[t[7]||(t[7]=l("strong",null,"正在播放：",-1)),l("span",{id:"current-song",ref_key:"currentSongDisplay",ref:w}," 未选择歌曲",512)]),l("div",ft,[o(d,{content:"上传音乐",effect:"light"},{default:a(()=>[o(u,{id:"upload-button",ref_key:"uploadButton",ref:H,onClick:t[2]||(t[2]=s=>v.value=!0),text:""},{default:a(()=>t[8]||(t[8]=[l("span",{class:"playing-icon"},"⏏️",-1)])),_:1},512)]),_:1}),o(d,{content:"播放方式",effect:"light"},{default:a(()=>[o(u,{id:"rand-button",ref:"randButton",onClick:z,text:""},{default:a(()=>t[9]||(t[9]=[l("span",{class:"playing-icon"},"🔀",-1)])),_:1},512)]),_:1}),o(d,{content:"上一首",effect:"light"},{default:a(()=>[o(u,{id:"prev-button",ref_key:"prevButton",ref:W,onClick:A,text:""},{default:a(()=>t[10]||(t[10]=[l("span",{class:"playing-icon"},"⬅️",-1)])),_:1},512)]),_:1}),o(d,{content:"播放/暂停",effect:"light"},{default:a(()=>[o(u,{id:"play-pause-button",onClick:j,ref_key:"playPauseButton",ref:q,text:""},{default:a(()=>[l("span",{class:"playing-icon",ref_key:"playPauseButtonText",ref:g},"▶️",512)]),_:1},512)]),_:1}),o(d,{content:"下一首",effect:"light"},{default:a(()=>[o(u,{id:"next-button",ref_key:"nextButton",ref:X,onClick:G,text:""},{default:a(()=>t[11]||(t[11]=[l("span",{class:"playing-icon"},"➡️",-1)])),_:1},512)]),_:1})]),l("div",null,[l("div",{class:"playing-time",id:"time-display",ref_key:"timeDisplay",ref:h},"0:00 / 0:00",512)])])])]),l("audio",{id:"audio-player",ref_key:"audioPlayer",ref:i,onTimeupdate:O,onEnded:Q},null,544),o(B,{modelValue:v.value,"onUpdate:modelValue":t[4]||(t[4]=s=>v.value=s),title:"上传音乐",width:"500","before-close":V},{footer:a(()=>[l("div",ct,[o(u,{onClick:V},{default:a(()=>t[13]||(t[13]=[L("Cancel")])),_:1}),o(u,{type:"primary",onClick:K},{default:a(()=>t[14]||(t[14]=[L(" Confirm ")])),_:1})])]),default:a(()=>[t[15]||(t[15]=l("span",null,"⚠️ 上传音乐时，文件名称应遵循格式：歌曲名 - 作者名。",-1)),t[16]||(t[16]=l("br",null,null,-1)),t[17]||(t[17]=l("br",null,null,-1)),o(x,{"file-list":m.value,"onUpdate:fileList":t[3]||(t[3]=s=>m.value=s),multiple:"","auto-upload":!1},{default:a(()=>[o(u,{type:"primary"},{default:a(()=>t[12]||(t[12]=[L("Click to upload")])),_:1})]),_:1},8,["file-list"])]),_:1},8,["modelValue"])],64)}}},kt=nt(pt,[["__scopeId","data-v-78c32018"]]);export{kt as default};