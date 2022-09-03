// 单个缓存条件设置
Vue.component("http-cache-ref-box", {
	props: ["v-cache-ref", "v-is-reverse"],
	mounted: function () {
		this.$refs.variablesDescriber.update(this.ref.key)
		if (this.ref.simpleCond != null) {
			this.changeCondType(this.ref.simpleCond.type)
			this.condCategory = "simple"
		} else if (this.ref.conds != null && this.ref.conds.groups != null) {
			this.condCategory = "complex"
		}
	},
	data: function () {
		let ref = this.vCacheRef
		if (ref == null) {
			ref = {
				isOn: true,
				cachePolicyId: 0,
				key: "${scheme}://${host}${requestPath}${isArgs}${args}",
				life: {count: 2, unit: "hour"},
				status: [200],
				maxSize: {count: 32, unit: "mb"},
				minSize: {count: 0, unit: "kb"},
				skipCacheControlValues: ["private", "no-cache", "no-store"],
				skipSetCookie: true,
				enableRequestCachePragma: false,
				conds: null, // 复杂条件
				simpleCond: null, // 简单条件
				allowChunkedEncoding: true,
				allowPartialContent: false,
				enableIfNoneMatch: false,
				enableIfModifiedSince: false,
				isReverse: this.vIsReverse,
				methods: [],
				expiresTime: {
					isPrior: false,
					isOn: false,
					overwrite: true,
					autoCalculate: true,
					duration: {count: -1, "unit": "hour"}
				}
			}
		}
		if (ref.key == null) {
			ref.key = ""
		}
		if (ref.methods == null) {
			ref.methods = []
		}

		if (ref.life == null) {
			ref.life = {count: 2, unit: "hour"}
		}
		if (ref.maxSize == null) {
			ref.maxSize = {count: 32, unit: "mb"}
		}
		if (ref.minSize == null) {
			ref.minSize = {count: 0, unit: "kb"}
		}

		let condType = "url-extension"
		let condComponent = window.REQUEST_COND_COMPONENTS.$find(function (k, v) {
			return v.type == "url-extension"
		})

		return {
			ref: ref,
			moreOptionsVisible: false,
			condCategory: "simple", // 条件分类：simple|complex
			condType: condType,
			condComponent: condComponent,
			components: window.REQUEST_COND_COMPONENTS
		}
	},
	methods: {
		changeOptionsVisible: function (v) {
			this.moreOptionsVisible = v
		},
		changeLife: function (v) {
			this.ref.life = v
		},
		changeMaxSize: function (v) {
			this.ref.maxSize = v
		},
		changeMinSize: function (v) {
			this.ref.minSize = v
		},
		changeConds: function (v) {
			this.ref.conds = v
			this.ref.simpleCond = null
		},
		changeStatusList: function (list) {
			let result = []
			list.forEach(function (status) {
				let statusNumber = parseInt(status)
				if (isNaN(statusNumber) || statusNumber < 100 || statusNumber > 999) {
					return
				}
				result.push(statusNumber)
			})
			this.ref.status = result
		},
		changeMethods: function (methods) {
			this.ref.methods = methods.map(function (v) {
				return v.toUpperCase()
			})
		},
		changeKey: function (key) {
			this.$refs.variablesDescriber.update(key)
		},
		changeExpiresTime: function (expiresTime) {
			this.ref.expiresTime = expiresTime
		},

		// 切换条件类型
		changeCondCategory: function (condCategory) {
			this.condCategory = condCategory
		},
		changeCondType: function (condType) {
			let def = this.components.$find(function (k, component) {
				return component.type == condType
			})
			if (def != null) {
				this.condComponent = def
			}
		}
	},
	template: `<tbody>
	<tr v-if="condCategory == 'simple'">
		<td class="title color-border">条件类型 *</td>
		<td>
			<select class="ui dropdown auto-width" name="condType" v-model="condType" @change="changeCondType(condType)">
				<option value="url-extension">URL扩展名</option>
				<option value="url-prefix">URL前缀</option>
				<option value="url-eq-index">首页</option>
				<option value="url-eq">URL精准匹配</option>
				<option value="url-regexp">URL正则匹配</option>
			</select>
			<p class="comment"><a href="" @click.prevent="changeCondCategory('complex')">切换到复杂条件 &raquo;</a></p>
		</td>
	</tr>
	<tr v-if="condCategory == 'simple'">
		<td class="color-border">{{condComponent.paramsTitle}} *</td>
		<td>
			<component :is="condComponent.component" :v-cond="ref.simpleCond"></component>
		</td>
	</tr>
	<tr v-if="condCategory == 'complex'">
		<td class="title">匹配条件分组 *</td>
		<td>
			<http-request-conds-box :v-conds="ref.conds" @change="changeConds"></http-request-conds-box>
			<p class="comment"><a href="" @click.prevent="changeCondCategory('simple')">&laquo; 切换到简单条件</a></p>
		</td>
	</tr>
	<tr v-show="!vIsReverse">
		<td>缓存有效期 *</td>
		<td>
			<time-duration-box :v-value="ref.life" @change="changeLife"></time-duration-box>
		</td>
	</tr>
	<tr v-show="!vIsReverse">
		<td>缓存Key *</td>
		<td>
			<input type="text" v-model="ref.key" @input="changeKey(ref.key)"/>
			<p class="comment">用来区分不同缓存内容的唯一Key。<request-variables-describer ref="variablesDescriber"></request-variables-describer>。</p>
		</td>
	</tr>
	<tr v-show="!vIsReverse">
		<td colspan="2"><more-options-indicator @change="changeOptionsVisible"></more-options-indicator></td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>请求方法限制</td>
		<td>
			<values-box size="5" maxlength="10" :values="ref.methods" @change="changeMethods"></values-box>
			<p class="comment">允许请求的缓存方法，默认支持所有的请求方法。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>客户端过期时间<em>（Expires）</em></td>
		<td>
			<http-expires-time-config-box :v-expires-time="ref.expiresTime" @change="changeExpiresTime"></http-expires-time-config-box>		
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>可缓存的最大内容尺寸</td>
		<td>
			<size-capacity-box :v-value="ref.maxSize" @change="changeMaxSize"></size-capacity-box>
			<p class="comment">内容尺寸如果高于此值则不缓存。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>可缓存的最小内容尺寸</td>
		<td>
			<size-capacity-box :v-value="ref.minSize" @change="changeMinSize"></size-capacity-box>
			<p class="comment">内容尺寸如果低于此值则不缓存。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>支持分片内容</td>
		<td>
			<checkbox name="allowChunkedEncoding" value="1" v-model="ref.allowChunkedEncoding"></checkbox>
			<p class="comment">选中后，Gzip等压缩后的Chunked内容可以直接缓存，无需检查内容长度。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>支持缓存区间内容</td>
		<td>
			<checkbox name="allowPartialContent" value="1" v-model="ref.allowPartialContent"></checkbox>
			<p class="comment">选中后，支持缓存源站返回的某个区间的内容，该内容通过<code-label>206 Partial Content</code-label>状态码返回。此功能目前为<code-label>试验性质</code-label>。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>状态码列表</td>
		<td>
			<values-box name="statusList" size="3" maxlength="3" :values="ref.status" @change="changeStatusList"></values-box>
			<p class="comment">允许缓存的HTTP状态码列表。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>跳过的Cache-Control值</td>
		<td>
			<values-box name="skipResponseCacheControlValues" size="10" maxlength="100" :values="ref.skipCacheControlValues"></values-box>
			<p class="comment">当响应的Cache-Control为这些值时不缓存响应内容，而且不区分大小写。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>跳过Set-Cookie</td>
		<td>
			<div class="ui checkbox">
				<input type="checkbox" value="1" v-model="ref.skipSetCookie"/>
				<label></label>
			</div>
			<p class="comment">选中后，当响应的Header中有Set-Cookie时不缓存响应内容。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>支持请求no-cache刷新</td>
		<td>
			<div class="ui checkbox">
				<input type="checkbox" name="enableRequestCachePragma" value="1" v-model="ref.enableRequestCachePragma"/>
				<label></label>
			</div>
			<p class="comment">选中后，当请求的Header中含有Pragma: no-cache或Cache-Control: no-cache时，会跳过缓存直接读取源内容。</p>
		</td>
	</tr>	
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>允许If-None-Match回源</td>
		<td>
			<checkbox v-model="ref.enableIfNoneMatch"></checkbox>
			<p class="comment">特殊情况下才需要开启，可能会降低缓存命中率。</p>
		</td>
	</tr>
	<tr v-show="moreOptionsVisible && !vIsReverse">
		<td>允许If-Modified-Since回源</td>
		<td>
			<checkbox v-model="ref.enableIfModifiedSince"></checkbox>
			<p class="comment">特殊情况下才需要开启，可能会降低缓存命中率。</p>
		</td>
	</tr>
	<tr v-show="false">
		<td colspan="2"><input type="hidden" name="cacheRefJSON" :value="JSON.stringify(ref)"/></td>
	</tr>
</tbody>`
})