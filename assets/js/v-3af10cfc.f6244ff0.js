"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[706],{1346:(e,i,t)=>{t.r(i),t.d(i,{data:()=>n});const n={key:"v-3af10cfc",path:"/intro/terminology.html",title:"Termonology",lang:"en-US",frontmatter:{},excerpt:"",headers:[{level:2,title:"LP Tokens",slug:"lp-tokens",children:[]},{level:2,title:"Liquidity Mining",slug:"liquidity-mining",children:[]},{level:2,title:"Impermanent Loss",slug:"impermanent-loss",children:[]},{level:2,title:"Long-Term Liquidity",slug:"long-term-liquidity",children:[]},{level:2,title:"IBC",slug:"ibc",children:[]}],filePathRelative:"intro/terminology.md",git:{updatedTime:1637530828e3,contributors:[{name:"Daniel Farina",email:"contact@nevulas.com",commits:1}]}}},7165:(e,i,t)=>{t.r(i),t.d(i,{default:()=>q});var n=t(6252);const s=(0,n.uE)('<h1 id="termonology" tabindex="-1"><a class="header-anchor" href="#termonology" aria-hidden="true">#</a> Termonology</h1><h2 id="lp-tokens" tabindex="-1"><a class="header-anchor" href="#lp-tokens" aria-hidden="true">#</a> LP Tokens</h2><p>When users deposit assets into a liquidity pool, they receive LP tokens. These tokens represent their share of the total pool. For example, if Pool #1 is the OSMO&lt;&gt;ATOM pool, users can deposit OSMO and ATOM tokens into the pool and receive back Pool1 share tokens. These tokens do not correspond to an exact quantity of tokens, but rather the proportional ownership of the pool. When users remove their liquidity from the pool, they get back the percentage of liquidity that their LP tokens represent. Since buying and selling from the pool changes the quantities of assets within a pool, users are highly unlikely to withdraw the same amount of each token that they initially deposited. They usually receive more of one and less of another, based on the trades executed from the pool.</p><h2 id="liquidity-mining" tabindex="-1"><a class="header-anchor" href="#liquidity-mining" aria-hidden="true">#</a> Liquidity Mining</h2>',4),o=(0,n.Uk)("Liquidity mining (also called yield farming) is when users earn tokens for providing liquidity to a DeFi protocol. This mechanism is used to offset the impermanent loss experienced by LPs. Liquidity mining rewards create an additional incentive for LPs besides transaction fees. These rewards are particularly useful for nascent protocols. Liquidity mining helps to bootstrap initial liquidity, facilitating increased usage and more fees for LPs. Information on Osmosis' incentive mining program can be found in this "),a={href:"https://osmosis.gitbook.io/o/osmo/token-issuance/liquidity-rewards",target:"_blank",rel:"noopener noreferrer"},r=(0,n.Uk)("section"),l=(0,n.Uk)(". [IMG1] [IMG2]"),h=(0,n._)("h2",{id:"impermanent-loss",tabindex:"-1"},[(0,n._)("a",{class:"header-anchor",href:"#impermanent-loss","aria-hidden":"true"},"#"),(0,n.Uk)(" Impermanent Loss")],-1),d=(0,n._)("p",null,"Liquidity providers earn through fees and special pool rewards. However, they are also risking a scenario in which they would have been better off holding the assets rather than supplying them. This outcome is called impermanent loss. Impermanent loss is the difference in net worth between HODLing and LPing. Liquidity mining helps to offset impermanent loss for LPs. When the price of the assets in the pool change at different rates, LPs end up owning larger amounts of the asset that increased less in price (or decreased more in price). For example, if the price of OSMO moons relative to ATOM, LPs in the OSMO-ATOM pool end up with larger portions of the less valuable asset (ATOM). [IMG3] Impermanent loss is mitigated in part by the transaction fees earned by LPs. When the profits made from swap fees outweigh an LP’s impermanent loss, the pool is self-sustainable. To further offset impermanent loss, particularly in the early stages of a protocol when volatility is high, AMMs utilize liquidity mining rewards. Liquidity rewards bootstrap the ecosystem as usage and fee revenues are still ramping up. Osmosis also has many new features and innovations in development to decrease impermanent loss as well.",-1),c=(0,n._)("h2",{id:"long-term-liquidity",tabindex:"-1"},[(0,n._)("a",{class:"header-anchor",href:"#long-term-liquidity","aria-hidden":"true"},"#"),(0,n.Uk)(" Long-Term Liquidity")],-1),u=(0,n.Uk)("Liquidity mining rewards tend to attract short-term “mercenary farmers” who quickly deposit and withdraw their liquidity after harvesting the yield. These farmers are only interested in the speculative value of the governance tokens that they are earning. They usually bounce between protocols in search of the best yield. Mercenary farmers often create the mirage of protocol adoption, but when these farmers leave, it results in significant liquidity volatility. Users of the AMM have difficulty executing trades without encountering slippage. Therefore, long-term liquidity is crucial to the success of an AMM. Osmosis’ design includes two mechanisms to incentivize long-term liquidity: "),m={href:"https://osmosis.gitbook.io/o/liquidity-providing/fees",target:"_blank",rel:"noopener noreferrer"},p=(0,n.Uk)("exit fees"),f=(0,n.Uk)(" and "),g={href:"https://osmosis.gitbook.io/o/liquidity-providing/blg",target:"_blank",rel:"noopener noreferrer"},y=(0,n.Uk)("bonded liquidity gauges"),b=(0,n.Uk)("."),k=(0,n._)("h2",{id:"ibc",tabindex:"-1"},[(0,n._)("a",{class:"header-anchor",href:"#ibc","aria-hidden":"true"},"#"),(0,n.Uk)(" IBC")],-1),w=(0,n._)("p",null,"The inter-blockchain communication protocol (IBC) creates communication between independent blockchains. IBC achieves this by specifying a set of structures that can be implemented by any distributed ledger that satisfies a small number of requirements. IBC facilitates cross-chain applications for token transfers, swaps, multi-chain contracts, and data sharding. At launch, Osmosis utilizes IBC for token transfers. Over time, Osmosis will add new features that are made possible through IBC.",-1),v={},q=(0,t(3744).Z)(v,[["render",function(e,i){const t=(0,n.up)("OutboundLink");return(0,n.wg)(),(0,n.iD)(n.HY,null,[s,(0,n._)("p",null,[o,(0,n._)("a",a,[r,(0,n.Wm)(t)]),l]),h,d,c,(0,n._)("p",null,[u,(0,n._)("a",m,[p,(0,n.Wm)(t)]),f,(0,n._)("a",g,[y,(0,n.Wm)(t)]),b]),k,w],64)}]])}}]);