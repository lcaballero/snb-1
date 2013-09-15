;(function($) {

	var defaults = {
		source: null,
		target: $("body:first"),
		url: null,
		data: {},
		template: "",
		url: null,
		append: true,
		replace: false
	};

	$.jadeDefaults = defaults;

	/**
	 * Extends the arguments provided but defaults the first argument
	 * to an empty object '{}' so that the first argument isn't 
	 * overriden, but instead the other properties are masked through
	 * the n-objects passed as parameters.
	 */
	function clone() {
		var args = [ {} ];
		args = args.concat( Array.prototype.slice.call(arguments) );

		var r = $.extend.apply( null, args );

		return r;
	}

	function requestJade(jq, opts) {

		opts = clone(defaults, opts);

		if (!!opts.template) {
			return Q.fcall(function() {
				return opts.template;
			});
		}

		if (!!opts.url) {
			return Q($.get(opts.url));
		}

		throw "Options didn't have a template or url.";
	};

	function renderTemplate(opts) {

		try {

			var tmpl = jade.compile(opts.template || "");
			var result = tmpl(opts.data || {});

		} catch(ex) {
			console.log("render exception", ex, opts)
		}

		if (opts.replace) {
			return opts.target.append($(result));
		} else {			
			return opts.target.append($(result));
		}
	};

	$.fn.promiseJade = function( opts ) {
		
		opts.target = this;

		return requestJade(this, opts)
			.then(function(v) {
				// Change a string selector to a jquery set
				if (typeof(v.source) == "string") {
					v.source = $(v.source);
				}
				return v;
			})
			.then(function(v) {
				var op = clone({ template:v }, opts);
				return op;
			})
			.then(renderTemplate);
	};

})(jQuery);











