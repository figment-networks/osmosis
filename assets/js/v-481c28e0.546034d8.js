"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[598],{6901:(s,n,a)=>{a.r(n),a.d(n,{data:()=>e});const e={key:"v-481c28e0",path:"/DOCS_README.html",title:"Osmosis Documentation Guide",lang:"en-US",frontmatter:{},excerpt:"",headers:[{level:3,title:"Vuepress",slug:"vuepress",children:[]},{level:3,title:"Config.js",slug:"config-js",children:[]},{level:3,title:"Local Development",slug:"local-development",children:[]},{level:2,title:"Production build & Automated Deployment to Github Pages",slug:"production-build-automated-deployment-to-github-pages",children:[]}],filePathRelative:"DOCS_README.md",git:{updatedTime:1637530828e3,contributors:[{name:"Daniel Farina",email:"contact@nevulas.com",commits:1}]}}},2493:(s,n,a)=>{a.r(n),a.d(n,{default:()=>C});var e=a(6252);const l=(0,e._)("h1",{id:"osmosis-documentation-guide",tabindex:"-1"},[(0,e._)("a",{class:"header-anchor",href:"#osmosis-documentation-guide","aria-hidden":"true"},"#"),(0,e.Uk)(" Osmosis Documentation Guide")],-1),o=(0,e._)("p",null,"The documentation for Osmosis is be hosted at:",-1),p={href:"https://docs.osmosis.zone/",target:"_blank",rel:"noopener noreferrer"},r=(0,e.Uk)("https://docs.osmosis.zone/"),t=(0,e.Uk)("This site is generated from files in this "),c={href:"https://github.com/osmosis-labs/osmosis/tree/master/docs",target:"_blank",rel:"noopener noreferrer"},i=(0,e._)("code",null,"docs",-1),d=(0,e.Uk)(" directory for "),u=(0,e._)("code",null,"master",-1),b=(0,e._)("h3",{id:"vuepress",tabindex:"-1"},[(0,e._)("a",{class:"header-anchor",href:"#vuepress","aria-hidden":"true"},"#"),(0,e.Uk)(" Vuepress")],-1),m=(0,e.Uk)("This site was generated using vuepress. "),h={href:"https://vuepress.vuejs.org/",target:"_blank",rel:"noopener noreferrer"},y=(0,e.Uk)("Documentation can be found here"),D=(0,e.Uk)("."),A=(0,e.uE)('<h3 id="config-js" tabindex="-1"><a class="header-anchor" href="#config-js" aria-hidden="true">#</a> Config.js</h3><p>The <a href="./.vuepress/config.js">config.js</a> contains most of the configuration used for the site generation.</p><h3 id="local-development" tabindex="-1"><a class="header-anchor" href="#local-development" aria-hidden="true">#</a> Local Development</h3><div class="language-text ext-text line-numbers-mode"><pre class="shiki" style="background-color:#22272e;"><code><span class="line"><span style="color:#adbac7;">git clone https://github.com/osmosis-labs/osmosis.git\n</span></span></code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>Make sure you are inside the docs folder</p><div class="language-text ext-text line-numbers-mode"><pre class="shiki" style="background-color:#22272e;"><code><span class="line"><span style="color:#adbac7;">cd docs\nyarn install\n</span></span></code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br></div></div><p>Run &amp; watch files for changes</p><div class="language-text ext-text line-numbers-mode"><pre class="shiki" style="background-color:#22272e;"><code><span class="line"><span style="color:#adbac7;">yarn dev\n</span></span></code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><h2 id="production-build-automated-deployment-to-github-pages" tabindex="-1"><a class="header-anchor" href="#production-build-automated-deployment-to-github-pages" aria-hidden="true">#</a> Production build &amp; Automated Deployment to Github Pages</h2><p>This site can be deployed to Github pages. All that needs to be done is use the git action included in the root directory under <code>.github/workflows/docbuild.yml</code>. This action will deploy the site inside a branch called gh-pages. This branch can then be configured to be served as a website in the repository settings. The git action looks like this:</p><div class="language-yaml ext-yml line-numbers-mode"><pre class="shiki" style="background-color:#22272e;"><code><span class="line"><span style="color:#8DDB8C;">name</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">Build and Deploy</span></span>\n<span class="line"><span style="color:#6CB6FF;">on</span><span style="color:#ADBAC7;">:</span></span>\n<span class="line"><span style="color:#ADBAC7;">  </span><span style="color:#8DDB8C;">push</span><span style="color:#ADBAC7;">:</span></span>\n<span class="line"><span style="color:#ADBAC7;">    </span><span style="color:#8DDB8C;">branches</span><span style="color:#ADBAC7;">:</span></span>\n<span class="line"><span style="color:#ADBAC7;">      - </span><span style="color:#96D0FF;">main</span></span>\n<span class="line"><span style="color:#8DDB8C;">jobs</span><span style="color:#ADBAC7;">:</span></span>\n<span class="line"><span style="color:#ADBAC7;">  </span><span style="color:#8DDB8C;">build-and-deploy</span><span style="color:#ADBAC7;">:</span></span>\n<span class="line"><span style="color:#ADBAC7;">    </span><span style="color:#8DDB8C;">runs-on</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">ubuntu-latest</span></span>\n<span class="line"><span style="color:#ADBAC7;">    </span><span style="color:#8DDB8C;">steps</span><span style="color:#ADBAC7;">:</span></span>\n<span class="line"><span style="color:#ADBAC7;">      - </span><span style="color:#8DDB8C;">name</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">Checkout 🛎️</span></span>\n<span class="line"><span style="color:#ADBAC7;">        </span><span style="color:#8DDB8C;">uses</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">actions/checkout@v2.3.1</span></span>\n<span class="line"></span>\n<span class="line"><span style="color:#ADBAC7;">      - </span><span style="color:#8DDB8C;">name</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">Install and Build 🔧</span><span style="color:#ADBAC7;"> </span><span style="color:#768390;"># This will create version inside the &#39;build&#39; folder.</span></span>\n<span class="line"><span style="color:#ADBAC7;">        </span><span style="color:#8DDB8C;">run</span><span style="color:#ADBAC7;">: </span><span style="color:#F47067;">|</span></span>\n<span class="line"><span style="color:#96D0FF;">          cd docs</span></span>\n<span class="line"><span style="color:#96D0FF;">          npm install</span></span>\n<span class="line"><span style="color:#96D0FF;">          npm run build</span></span>\n<span class="line"></span>\n<span class="line"><span style="color:#ADBAC7;">      - </span><span style="color:#8DDB8C;">name</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">Deploy 🚀</span></span>\n<span class="line"><span style="color:#ADBAC7;">        </span><span style="color:#8DDB8C;">uses</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">JamesIves/github-pages-deploy-action@4.1.5</span></span>\n<span class="line"><span style="color:#ADBAC7;">        </span><span style="color:#8DDB8C;">with</span><span style="color:#ADBAC7;">:</span></span>\n<span class="line"><span style="color:#ADBAC7;">          </span><span style="color:#8DDB8C;">branch</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">gh-pages</span><span style="color:#ADBAC7;"> </span><span style="color:#768390;"># The branch the action should deploy to.</span></span>\n<span class="line"><span style="color:#ADBAC7;">          </span><span style="color:#8DDB8C;">folder</span><span style="color:#ADBAC7;">: </span><span style="color:#96D0FF;">docs/src/.vuepress/dist</span></span>\n<span class="line"></span></code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br></div></div>',11),g={},C=(0,a(3744).Z)(g,[["render",function(s,n){const a=(0,e.up)("OutboundLink");return(0,e.wg)(),(0,e.iD)(e.HY,null,[l,o,(0,e._)("ul",null,[(0,e._)("li",null,[(0,e._)("a",p,[r,(0,e.Wm)(a)])])]),(0,e._)("p",null,[t,(0,e._)("a",c,[i,d,u,(0,e.Wm)(a)])]),b,(0,e._)("p",null,[m,(0,e._)("a",h,[y,(0,e.Wm)(a)]),D]),A],64)}]])}}]);