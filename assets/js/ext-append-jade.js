;(function($) {

	function renderTemplate(opts) {

		var tmpl = jade.compile(opts.template || "");
		var result = tmpl(opts.data || {});

		if (opts.replace) {
			return this.append(result);
		} else {
			return this.append(result);
		}
	};

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

	/**
	 * Append to the jQuery selection the rendered template that
	 * can be found as html inside of the element provided as
	 * @template.  @template can be a selector, a jquery seleciton,
	 * a DOM element, basically anything that can be passed to $().
	 *
	 * The @data paramter is passed to the template during template
	 * compilation.
	 *
	 * {
	 *	source: [ "selector", $(), DOM ],
	 *  target: [ this ]
	 *	data: {},
	 *	template: "",
	 *	url: "/url/to/a/.jade/file",
	 *  append: true,
	 *  replace: false
	 * }
	 */
	$.fn.appendJade = function(opts) {

		opts = $.extend({}, defaults, opts);

		// Without one of these there's nothing to consider as Jade
		// and then render to the DOM
		if (!opts.source && !opts.url && !opts.template) {
			return this;
		}

		// Change a string selector to a jquery set
		if (typeof(opts.source) == "string") {
			opts.source = $(opts.source);
		}
		
		if (!!opts.url) {
			console.log("requestiong jade:", opts.url)
			var self = this;
			$.ajax({
				type:"GET",
				url:opts.url,
				success:function(r) {
					opts.template = r;
					console.log("requested jade: ", opts);
					renderTemplate.apply(self, [ opts ]);
				}
			});
		} else {
			opts.template = opts.source.html();
			return renderTemplate.apply(this, [ opts ]);
		}
	};

})(jQuery);