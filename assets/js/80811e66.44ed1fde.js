"use strict";(self.webpackChunkdocumentation=self.webpackChunkdocumentation||[]).push([[616],{3905:(e,t,i)=>{i.d(t,{Zo:()=>c,kt:()=>m});var n=i(7294);function p(e,t,i){return t in e?Object.defineProperty(e,t,{value:i,enumerable:!0,configurable:!0,writable:!0}):e[t]=i,e}function r(e,t){var i=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),i.push.apply(i,n)}return i}function a(e){for(var t=1;t<arguments.length;t++){var i=null!=arguments[t]?arguments[t]:{};t%2?r(Object(i),!0).forEach((function(t){p(e,t,i[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(i)):r(Object(i)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(i,t))}))}return e}function o(e,t){if(null==e)return{};var i,n,p=function(e,t){if(null==e)return{};var i,n,p={},r=Object.keys(e);for(n=0;n<r.length;n++)i=r[n],t.indexOf(i)>=0||(p[i]=e[i]);return p}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(n=0;n<r.length;n++)i=r[n],t.indexOf(i)>=0||Object.prototype.propertyIsEnumerable.call(e,i)&&(p[i]=e[i])}return p}var s=n.createContext({}),l=function(e){var t=n.useContext(s),i=t;return e&&(i="function"==typeof e?e(t):a(a({},t),e)),i},c=function(e){var t=l(e.components);return n.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var i=e.components,p=e.mdxType,r=e.originalType,s=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),u=l(i),m=p,f=u["".concat(s,".").concat(m)]||u[m]||d[m]||r;return i?n.createElement(f,a(a({ref:t},c),{},{components:i})):n.createElement(f,a({ref:t},c))}));function m(e,t){var i=arguments,p=t&&t.mdxType;if("string"==typeof e||p){var r=i.length,a=new Array(r);a[0]=u;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o.mdxType="string"==typeof e?e:p,a[1]=o;for(var l=2;l<r;l++)a[l]=i[l];return n.createElement.apply(null,a)}return n.createElement.apply(null,i)}u.displayName="MDXCreateElement"},4411:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>s,contentTitle:()=>a,default:()=>d,frontMatter:()=>r,metadata:()=>o,toc:()=>l});var n=i(7462),p=(i(7294),i(3905));const r={id:"apps",title:"apps",hide_title:!0,sidebar_label:"apps",description:"Application commands",last_update:{date:new Date("2022-11-24T20:58:11.000Z"),author:"blacktop"}},a=void 0,o={unversionedId:"cli/ipsw/idev/apps/apps",id:"cli/ipsw/idev/apps/apps",title:"apps",description:"Application commands",source:"@site/docs/cli/ipsw/idev/apps/apps.md",sourceDirName:"cli/ipsw/idev/apps",slug:"/cli/ipsw/idev/apps/",permalink:"/ipsw/docs/cli/ipsw/idev/apps/",draft:!1,editUrl:"https://github.com/blacktop/ipsw/tree/master/www/docs/cli/ipsw/idev/apps/apps.md",tags:[],version:"current",frontMatter:{id:"apps",title:"apps",hide_title:!0,sidebar_label:"apps",description:"Application commands",last_update:{date:"2022-11-24T20:58:11.000Z",author:"blacktop"}},sidebar:"cli",previous:{title:"tree",permalink:"/ipsw/docs/cli/ipsw/idev/afc/tree"},next:{title:"install",permalink:"/ipsw/docs/cli/ipsw/idev/apps/install"}},s={},l=[{value:"ipsw idev apps",id:"ipsw-idev-apps",level:2},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],c={toc:l};function d(e){let{components:t,...i}=e;return(0,p.kt)("wrapper",(0,n.Z)({},c,i,{components:t,mdxType:"MDXLayout"}),(0,p.kt)("h2",{id:"ipsw-idev-apps"},"ipsw idev apps"),(0,p.kt)("p",null,"Application commands"),(0,p.kt)("pre",null,(0,p.kt)("code",{parentName:"pre"},"ipsw idev apps [flags]\n")),(0,p.kt)("h3",{id:"options"},"Options"),(0,p.kt)("pre",null,(0,p.kt)("code",{parentName:"pre"},"  -h, --help   help for apps\n")),(0,p.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,p.kt)("pre",null,(0,p.kt)("code",{parentName:"pre"},"      --color           colorize output\n      --config string   config file (default is $HOME/.ipsw.yaml)\n  -u, --udid string     Device UniqueDeviceID to connect to\n  -V, --verbose         verbose output\n")),(0,p.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,p.kt)("ul",null,(0,p.kt)("li",{parentName:"ul"},(0,p.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/idev"},"ipsw idev"),"\t - USB connected device commands"),(0,p.kt)("li",{parentName:"ul"},(0,p.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/idev/apps/install"},"ipsw idev apps install"),"\t - Install an application"),(0,p.kt)("li",{parentName:"ul"},(0,p.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/idev/apps/ls"},"ipsw idev apps ls"),"\t - List installed applications"),(0,p.kt)("li",{parentName:"ul"},(0,p.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/idev/apps/uninstall"},"ipsw idev apps uninstall"),"\t - Uninstall an application")))}d.isMDXComponent=!0}}]);