"use strict";(self.webpackChunkdocumentation=self.webpackChunkdocumentation||[]).push([[6603],{3905:(e,t,i)=>{i.d(t,{Zo:()=>c,kt:()=>f});var n=i(67294);function r(e,t,i){return t in e?Object.defineProperty(e,t,{value:i,enumerable:!0,configurable:!0,writable:!0}):e[t]=i,e}function o(e,t){var i=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),i.push.apply(i,n)}return i}function l(e){for(var t=1;t<arguments.length;t++){var i=null!=arguments[t]?arguments[t]:{};t%2?o(Object(i),!0).forEach((function(t){r(e,t,i[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(i)):o(Object(i)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(i,t))}))}return e}function p(e,t){if(null==e)return{};var i,n,r=function(e,t){if(null==e)return{};var i,n,r={},o=Object.keys(e);for(n=0;n<o.length;n++)i=o[n],t.indexOf(i)>=0||(r[i]=e[i]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)i=o[n],t.indexOf(i)>=0||Object.prototype.propertyIsEnumerable.call(e,i)&&(r[i]=e[i])}return r}var a=n.createContext({}),s=function(e){var t=n.useContext(a),i=t;return e&&(i="function"==typeof e?e(t):l(l({},t),e)),i},c=function(e){var t=s(e.components);return n.createElement(a.Provider,{value:t},e.children)},d="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},v=n.forwardRef((function(e,t){var i=e.components,r=e.mdxType,o=e.originalType,a=e.parentName,c=p(e,["components","mdxType","originalType","parentName"]),d=s(i),v=r,f=d["".concat(a,".").concat(v)]||d[v]||u[v]||o;return i?n.createElement(f,l(l({ref:t},c),{},{components:i})):n.createElement(f,l({ref:t},c))}));function f(e,t){var i=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=i.length,l=new Array(o);l[0]=v;var p={};for(var a in t)hasOwnProperty.call(t,a)&&(p[a]=t[a]);p.originalType=e,p[d]="string"==typeof e?e:r,l[1]=p;for(var s=2;s<o;s++)l[s]=i[s];return n.createElement.apply(null,l)}return n.createElement.apply(null,i)}v.displayName="MDXCreateElement"},7651:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>a,contentTitle:()=>l,default:()=>u,frontMatter:()=>o,metadata:()=>p,toc:()=>s});var n=i(87462),r=(i(67294),i(3905));const o={id:"install",title:"install",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"install",description:"Install a provision profile (.mobileprovision file)"},l=void 0,p={unversionedId:"cli/ipsw/idev/prov/install",id:"cli/ipsw/idev/prov/install",title:"install",description:"Install a provision profile (.mobileprovision file)",source:"@site/docs/cli/ipsw/idev/prov/install.md",sourceDirName:"cli/ipsw/idev/prov",slug:"/cli/ipsw/idev/prov/install",permalink:"/ipsw/docs/cli/ipsw/idev/prov/install",draft:!1,editUrl:"https://github.com/blacktop/ipsw/tree/master/www/docs/cli/ipsw/idev/prov/install.md",tags:[],version:"current",frontMatter:{id:"install",title:"install",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"install",description:"Install a provision profile (.mobileprovision file)"},sidebar:"cli",previous:{title:"dump",permalink:"/ipsw/docs/cli/ipsw/idev/prov/dump"},next:{title:"ls",permalink:"/ipsw/docs/cli/ipsw/idev/prov/ls"}},a={},s=[{value:"ipsw idev prov install",id:"ipsw-idev-prov-install",level:2},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],c={toc:s},d="wrapper";function u(e){let{components:t,...i}=e;return(0,r.kt)(d,(0,n.Z)({},c,i,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h2",{id:"ipsw-idev-prov-install"},"ipsw idev prov install"),(0,r.kt)("p",null,"Install a provision profile (.mobileprovision file)"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"ipsw idev prov install <PROV> [flags]\n")),(0,r.kt)("h3",{id:"options"},"Options"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"  -h, --help   help for install\n")),(0,r.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"      --color           colorize output\n      --config string   config file (default is $HOME/.ipsw/config.yaml)\n  -u, --udid string     Device UniqueDeviceID to connect to\n  -V, --verbose         verbose output\n")),(0,r.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/idev/prov"},"ipsw idev prov"),"\t - Provision commands")))}u.isMDXComponent=!0}}]);