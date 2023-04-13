"use strict";(self.webpackChunkdocumentation=self.webpackChunkdocumentation||[]).push([[7167],{3905:(e,t,a)=>{a.d(t,{Zo:()=>p,kt:()=>u});var r=a(67294);function o(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function i(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,r)}return a}function n(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?i(Object(a),!0).forEach((function(t){o(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):i(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function c(e,t){if(null==e)return{};var a,r,o=function(e,t){if(null==e)return{};var a,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)a=i[r],t.indexOf(a)>=0||(o[a]=e[a]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)a=i[r],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(o[a]=e[a])}return o}var l=r.createContext({}),s=function(e){var t=r.useContext(l),a=t;return e&&(a="function"==typeof e?e(t):n(n({},t),e)),a},p=function(e){var t=s(e.components);return r.createElement(l.Provider,{value:t},e.children)},m="mdxType",h={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var a=e.components,o=e.mdxType,i=e.originalType,l=e.parentName,p=c(e,["components","mdxType","originalType","parentName"]),m=s(a),d=o,u=m["".concat(l,".").concat(d)]||m[d]||h[d]||i;return a?r.createElement(u,n(n({ref:t},p),{},{components:a})):r.createElement(u,n({ref:t},p))}));function u(e,t){var a=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=a.length,n=new Array(i);n[0]=d;var c={};for(var l in t)hasOwnProperty.call(t,l)&&(c[l]=t[l]);c.originalType=e,c[m]="string"==typeof e?e:o,n[1]=c;for(var s=2;s<i;s++)n[s]=a[s];return r.createElement.apply(null,n)}return r.createElement.apply(null,a)}d.displayName="MDXCreateElement"},63925:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>l,contentTitle:()=>n,default:()=>h,frontMatter:()=>i,metadata:()=>c,toc:()=>s});var r=a(87462),o=(a(67294),a(3905));const i={id:"macho",title:"macho",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"macho",description:"Parse MachO"},n=void 0,c={unversionedId:"cli/ipsw/macho/macho",id:"cli/ipsw/macho/macho",title:"macho",description:"Parse MachO",source:"@site/docs/cli/ipsw/macho/macho.md",sourceDirName:"cli/ipsw/macho",slug:"/cli/ipsw/macho/",permalink:"/ipsw/docs/cli/ipsw/macho/",draft:!1,editUrl:"https://github.com/blacktop/ipsw/tree/master/www/docs/cli/ipsw/macho/macho.md",tags:[],version:"current",frontMatter:{id:"macho",title:"macho",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"macho",description:"Parse MachO"},sidebar:"cli",previous:{title:"version",permalink:"/ipsw/docs/cli/ipsw/kernel/version"},next:{title:"a2o",permalink:"/ipsw/docs/cli/ipsw/macho/a2o"}},l={},s=[{value:"ipsw macho",id:"ipsw-macho",level:2},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],p={toc:s},m="wrapper";function h(e){let{components:t,...a}=e;return(0,o.kt)(m,(0,r.Z)({},p,a,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"ipsw-macho"},"ipsw macho"),(0,o.kt)("p",null,"Parse MachO"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"ipsw macho [flags]\n")),(0,o.kt)("h3",{id:"options"},"Options"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"  -h, --help   help for macho\n")),(0,o.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"      --color           colorize output\n      --config string   config file (default is $HOME/.ipsw/config.yaml)\n  -V, --verbose         verbose output\n")),(0,o.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw"},"ipsw"),"\t - Download and Parse IPSWs (and SO much more)"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/a2o"},"ipsw macho a2o"),"\t - Convert MachO address to offset"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/a2s"},"ipsw macho a2s"),"\t - Lookup symbol at unslid address"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/bbl"},"ipsw macho bbl"),"\t - Create single universal/fat MachO out many MachOs"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/disass"},"ipsw macho disass"),"\t - Disassemble ARM64 MachO at symbol/vaddr"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/dump"},"ipsw macho dump"),"\t - Dump MachO data at given virtual address"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/info"},"ipsw macho info"),"\t - Explore a MachO file"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/lipo"},"ipsw macho lipo"),"\t - Extract single MachO out of a universal/fat MachO"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/o2a"},"ipsw macho o2a"),"\t - Convert MachO offset to address"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/patch"},"ipsw macho patch"),"\t - Patch MachO Load Commands"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/search"},"ipsw macho search"),"\t - Find Mach-O files for given search criteria"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/macho/sign"},"ipsw macho sign"),"\t - Codesign a MachO")))}h.isMDXComponent=!0}}]);