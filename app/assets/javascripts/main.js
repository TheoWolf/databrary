// clean dbjs namespace
var dbjs = {};

dbjs.vars = {
	speedFast: 150,
	speedNorm: 250,
	speedSlow: 350,
	speedTest: 1000
};

dbjs.baselineFigure = function (elements) {
	var $elements = $(elements),
		sizes = {
			square: 1,
			wide: .5625,
			tube: .75
		},
		mult = 12;

	var resize = function () {
		$elements.each(function (e) {
			var $el = $(this),
				width = $el.outerWidth(false),
				size = sizes['square'],
				height;

			for (var key in sizes) {
				if (sizes.hasOwnProperty(key) && $el.hasClass(key)) {
					size = sizes[key];
					break;
				}
			}

			height = width * size;
			height = height - (height % mult);

			$el.height(height);
		});
	};

	var timeout;

	$(window).resize(function () {
		clearTimeout(timeout);
		timeout = setTimeout(function () {
			resize();
		}, 500);
	});

	resize();
};

dbjs.hoverable = function (elements) {
	var $body = $('body'),
		$els = $(elements),
		$wrap = $('<div class="hoverable_wrap"></div>');

	$els.wrap($wrap);

	$body.on('mouseenter mouseleave', elements, function () {
		var $el = $(this),
			$wrap = $el.parent(),
			timeout;

		if($el.is(':hover')){
			clearTimeout(timeout);

			$el.css('z-index', 750);
			$wrap.stop().height($el.outerHeight(true)).width($el.outerWidth(true));
			$el.addClass('hoverable_float');
		}else{
			clearTimeout(timeout);

			timeout = setTimeout(function () {
				$wrap.stop().css({
					'width': '',
					'height': ''
				});

				$el.removeClass('hoverable_float');
			}, 150);

			$el.css('z-index', 0);
		}
	});
};

dbjs.minimizable = function (elements) {
	var $els = $(elements);

	$els.find('.minimizer').click(function() {
		$(this).closest(elements).toggleClass('minimized').find('.minimize').slideToggle(500);
	})
};

/**
 * Creates tabset with content panes
 * @param tabset    containing element
 * @param tab        class of tabs
 * @param body        class of content panes
 */
dbjs.tabs = function (tabset, tab, body) {
	var $tabset = $(tabset),
		$tabs = $tabset.find(tab),
		$bodies = $tabset.find(body),
		$tabs_el = $('<ul class="tabs cf"></ul>'),
		$body_el = $('<div class="views cf"></div>');

	// rebuild dom
	$tabset.prepend($tabs_el);
	$tabset.append($body_el);

	$tabs.each(function () {
		var $this = $(this),
			id = $this.parent().attr('id'),
			tip = $this.attr('data-message');

		$this.appendTo($tabs_el).replaceWith($('<li class="tab"><a href="#' + id + '" data-message="' + tip + '">' + this.innerHTML + '</a></li>'));
	});

	if ($messageHandler)
		$messageHandler.data('messageHandler').generate(tabset + ' .tab > a', 'data-message', 'trace');

	$bodies.each(function () {
		$(this).appendTo($body_el).addClass('rolled').slideUp(0);
	});

	// tab clicking
	$tabs = $('ul.tabs ' + tab);
	$bodies = $('.views ' + body);

	var $tablinks = $tabs.find('a'),
		$currentBody = $bodies.first();

	var clickTab = function (tab) {
		var $this = $(tab),
			id = $this.attr('href'),
			$body = $(id);

		if ($this.parent().hasClass('current')) {
			$currentBody.fadeOut(150, function () {
				$currentBody.fadeIn(150);
			});

			return;
		}

		$currentBody.slideUp(250, function () {
			$body.slideDown(250);
		});

		$tabs.removeClass('current');
		$this.parent().addClass('current');

		$currentBody = $body;
	};

	$tablinks.click(function () {
		clickTab(this);

		return false;
	});

	clickTab($tablinks.first());
};

/**
 * Creates rolldown modal content.
 * @param clicker    click toggle
 * @param toggle    toggled area
 * @param now        display modal now, or on click
 */
dbjs.modal = function (clicker, toggle, now) {
	if (typeof(now) === 'undefined') now = false;

	var $clicker = $(clicker),
		$toggle = $(toggle),
		modals = '.modal',
		back = '#modal_back',
		$back,
		speed = 250;

	// modal background
	if (!$(back).exists()) {
		$('body').append('<div id="modal_back"></div>');
		$back = $(back);

		$back.click(function () {
			$(modals).slideUp(speed);
			$(this).fadeOut(speed);
		});
	}

	// clicker
	var click = function () {
		$toggle.slideDown(speed);
		$back.fadeIn(speed);
		$("html, body").animate({ scrollTop: 0 }, speed);
	};

	$clicker.click(function () {
		click();
		return false;
	});

	if (now === true) {
		click();
	}
};

/**
 * Creates rolldown modal content filled by ajax.
 * @param clicker    click toggle
 * @param url        ajax retrieved content
 * @param now        load content now, or on click
 */
dbjs.ajaxModal = function (clicker, url, now) {
	if (typeof(now) === 'undefined') now = false;

	var $clicker = $(clicker),
		$toggle, toggle;

	if (!$clicker.exists())
		return;

	var setup = function (url) {
		$clicker.off('click');

		$.get(url, function (data) {
			var $data = $(data),
				$script = $data.filter('script');

			$toggle = $data.find('.modal');
			toggle = $clicker.attr('data-target');

			$toggle.attr('id', toggle).appendTo($('#site_content'));

			toggle = '#' + toggle;

			if ($script.exists())
				$script.appendTo($('body'));

			if (now === true) {
				dbjs.modal(clicker, toggle, false);
			} else {
				$clicker.off('click.ajaxModal');
				dbjs.modal(clicker, toggle, true);
			}

		}, 'html');

		$clicker.on('click');
	};

	if (now === true) {
		setup(url);
	} else {
		$clicker.on('click.ajaxModal', function () {
			setup(url);
		});
	}
};

/**
 * Creates folding content.
 * @param region    container element
 * @param clicker    click toggle, child of fold
 * @param toggle    toggled area, child of fold
 */
dbjs.fold = function (region, clicker, toggle) {
	var $region = $(region),
		$clicker = $region.find(clicker),
		$toggles = $region.find(toggle);

	$clicker.click(function () {
		var $this = $(this);

		$this.closest(region).find(toggle).slideToggle();
		$this.toggleClass('unfolded');
	});

	$clicker.prepend('<span class="arrow"></span>');
	$toggles.toggle();
};

dbjs.carousel = function (region, panel) {
	var $region = $(region),
		$panels = $region.find(panel);

	var interval = setInterval(function (e) {
		$region.find(">" + panel + ":last-child").fadeOut(1000, function () {
			$(this).prependTo($region).fadeIn(1000);
		});
	}, 5000);
};

/**
 * Creates a compact menu for small browser windows
 * @param menu        region to resize
 * @param position    where to position menu link
 */
dbjs.sideMenu = function (menu, position) {
	var $menu = $(menu),
		$position = $(position),
		$origin = $menu.parent(),
		$menuHeader = $menu.find('h1'),
		$menuLink = $('<button class="hide"></button>'),
		toggle = false;

	// build menu elements
	$menuLink.prependTo($menu);

	// hide and show based on browser size and zoom
	var reducer = function () {
		if (toggle == false && $menu.css('z-index') == 1) {
			$menu.addClass('reduced').prependTo($position);
			$menuHeader.remove();
			$menuLink.removeClass('hide');
			toggle = true;
		} else if (toggle == true && $menu.css('z-index') == 'auto') {
			$menu.removeClass('reduced').prependTo($origin);
			$menuHeader.prependTo($menu);
			$menuLink.addClass('hide');
			toggle = false;
		}
	};

	$(window).resize(function () {
		reducer();
	});

	reducer();

	// make menu clickable
	$menuLink.click(function () {
		$menu.toggleClass('dropped');
	});

	$menu.click(function (e) {
		e.stopPropagation();
	});

	$(document).click(function () {
		$menu.removeClass('dropped');
	});
};

/**
 * Keeps footer at the bottom of the window when the content is shorter than the window
 * @param footer    element made sticky
 * @param above    content to monitor for height
 */
dbjs.stickyFooter = function (footer, above) {
	var $above = $(above),
		aboveW = $above.width(),
		aboveH = $above.height();

	// monitor window
	var resize = function (footer, above) {
		var windowHeight = $(window).height(),
			$footer = $(footer),
			$above = $(above),
			footerOffset = $footer.outerHeight() + $above.outerHeight(true);

		if (windowHeight > footerOffset)
			$footer.addClass('fixed');
		else
			$footer.removeClass('fixed');
	};

	$(window).resize(function () {
		resize(footer, above);
	});

	// monitor content
	var checkResize = function () {
		var w = $above.width(),
			h = $above.height();

		if (w != aboveW || h != aboveH) {
			aboveW = w;
			aboveH = h;

			return true;
		}

		return false;
	};

	var contentCheck = setInterval(function () {
		if (checkResize())
			resize(footer, above);
	}, 50);

	resize(footer, above);
};

/**
 * Fades the bottom of longer content based on shorter content size
 * @param container wrapper for fader and faded element
 * @param faded    element faded to fit
 * @param fader    element measured for height
 */
dbjs.fadeOff = function (container, faded, fader) {
	var $containers = $(container),
		fade = '<div class="fade"></div>',
		$imgs = $containers.find(fader + ' img');

	// $img.load().each(if this.complete, load...) doesn't work in modal.
	// ridiculous workaround with height minimum.
	var loaded = 0;

	$imgs.load(function () {
		loaded++;

		if ($imgs.length == loaded) {
			$containers.each(function () {
				var $this = $(this),
					$faded = $this.find(faded),
					$fader = $this.find(fader);

				if ($fader.height() > 150)
					$faded.height($fader.height());
				else
					$faded.height(150);

				$faded.append($(fade));
			});
		}
	});
};

/**
 * Uses a set of togglers to hide all but the active element and the inverse set of togglers
 * @param toggler    elements used to toggle
 * @param toggled    elements toggled by togglers
 */
dbjs.simpleToggle = function (toggler, toggled) {
	var $togglers = $(toggler),
		$toggleds = $(toggled),
		$active = $togglers.filter('.active');

	var clicker = function (toggler) {
		var $toggler = $(toggler),
			$toggled = $($toggler.attr('href'));

		$toggler.fadeOut(dbjs.vars.speedNorm, function () {
			$togglers.not($toggler).fadeIn(dbjs.vars.speedNorm);
		});

		$toggleds.not($toggled).slideUp(dbjs.vars.speedNorm, function () {
			$toggled.slideDown(dbjs.vars.speedNorm);
		});
	};

	$togglers.click(function (e) {
		e.stopPropagation();

		clicker(this);
	});

	if ($active.exists())
		clicker($active);
};

/**
 * MESSAGES AND MESSAGE HANDLER
 */
(function ($, window, document) {
	// stock vars
	var $window = $(window),
		$handler,
		sources = ['auto', 'alt', 'title', 'html', 'data-message', 'data-source'],
		types = ['alert', 'error', 'trace', 'input', 'null'];

	var validMessage = function (message) {
		return $.type(message) == 'string'
	};

	var validType = function (type) {
		return $.inArray(type, types) >= 0;
	};

	var defaultType = function ($source, type) {
		if ($.inArray(type, types) < 0)
			type = $source.attr('class').split(/\s+/).filter(function (n) {
				return types.indexOf(n) != -1;
			}).shift();

		if (!$.inArray(type, types) < 0)
			type = 'trace';

		return type;
	};

	$.extend($.fn, {
		/**
		 *
		 * @param args
		 * @returns {*}
		 */
		messageHandler: function (args) {
			//options
			args = $.extend({
				speed: 150
			}, args);

			$handler = this;

			var messages = {},
				$alerts = $('<div class="alerts"></div>').appendTo($handler),
				$errors = $('<div class="errors"></div>').appendTo($handler),
				$inputs = $('<div class="inputs"></div>').appendTo($handler),
				$traces = $('<div class="traces"></div>').appendTo($handler);

			var router = {
				alert: $alerts,
				error: $errors,
				trace: $traces,
				input: $inputs
			};

			// methods
			var initialize = function () {
				generate('.alert', 'html', 'alert');
				generate('#site_header .message', 'data-message', 'trace');
				generate('.thumb', 'data-message', 'trace');
				generate('.message-trace', 'data-message', 'trace');
				//generate('.error', 'html', 'error');
			};

			var generate = function (element, source, type, prepend) {
				var $elements = $(element),
					messages = {};

				$elements.each(function () {
					var $element = $(this),
						$message;

					if ($.inArray(source, sources) < 0 && !$(source).exists()) source = 'auto';
					type = defaultType($element, type);

					$message = getMessage($element, source, type);

					if ($message && create($element, $message, type, args, prepend))
						messages[$message.attr('id')] = $message;
				});

				return messages;
			};

			var getMessage = function ($element, source, type) {
				var messageID = 'message-' + new Date().getTime().toString().substr(-5) + Math.floor((Math.random() * 1000) + 1),
					$message;

				// pick 'auto' source
				if (source == 'auto') {
					if ($element.attr('data-source') != '')
						source = 'data-source';
					else if ($element.attr('data-message') != '')
						source = 'data-message';
					else if ($element.attr('title') != '')
						source = 'title';
					else if ($element.attr('alt') != '')
						source = 'alt';
					else // no source? that's dumb.
						return false;
				}

				// grab the content
				if (source == 'data-source')
					$message = $($element.attr('data-source'));
				else if (source == 'html')
					$message = $('<div></div>').html($element.remove().html());
				else if ($.inArray(source, sources) >= 0)
					$message = $('<div></div>').html($element.attr(source));
				else
					$message = $(source);

				if ($message.text() == '')
					return false;

				$message.attr('id', messageID).addClass('message');

				// clean up
				$element
					.addClass('toggle_message toggle_' + type)
					.attr('data-source', '#' + messageID)
					.removeAttr('alt')
					.removeAttr('title')
					.removeAttr('data-message');

				return $message;
			};

			var create = function ($element, $message, type, args, prepend) {
				type = defaultType($message, type);
				$message = $message.messageHelper(type, args);

				if (!$message)
					return false;

				move($element, $message, type, prepend);

				addHook($element, $message, type);

				return $message;
			};

			var update = function ($element, message, type, args, prepend) {
				var $message = $($element.attr('data-source'));

				if (!$message || $handler.children($message) == 0)
					return false;

				if (!$message.hasClass(type)) {
					removeHook($element, $message, $message.attr('class').split(/\s+/).filter(function (n) {
						return types.indexOf(n) != -1;
					}).shift());
					addHook($element, $message, type);
				}

				if (!$message.data('messageHelper').update(message, type, args))
					return false;

				move($element, $message, type, prepend);

				return $message
			};

			var move = function ($element, $message, type, prepend) {
				if (prepend !== true)
					router[type].append($message);
				else
					router[type].prepend($message);

				reorderMessages(type);

				return $message;
			};

			var reorderMessages = function (type) {
				var $messages = router[type].find('.message'),
					messageCount = $messages.length;

				$messages.each(function (i, e) {
					$(e).css('z-index', messageCount - i);
				});
			};

			var addHook = function ($element, $message, type) {
				switch (type) {
					case 'trace':
						$element.hover(function () {
							show($message);
						}, function () {
							hide($message);
						});

						hide($message);

						break;
					case 'alert':
						show($message, null, true);

						addCloser($message);

						break;
					case 'input':
					case 'error':
						$element.focusin(function () {
							show($message);
						});

						$element.focusout(function () {
							hide($message);
						});

						hide($message);

						break;
				}
			};

			var addCloser = function ($message) {
				$('<div class="x"></div>').on('click',function () {
					hide($message, null, true);
				}).prependTo($message.find('.wrap'));
			};

			var removeHook = function ($element, $message, type) {
				switch (type) {
					case 'trace':
						$element.off('hover');

						hide($message);

						break;
					case 'alert':
						hide($message, null, true);

						removeCloser($message);

						break;
					case 'input':
					case 'error':
						$element.off('focusin');

						$element.off('focusout');

						hide($message);

						break;
				}
			};

			var removeCloser = function ($message) {
				$message.find('.close').off('click').remove();
			};

			var remove = function ($message, callback, adjust) {
				return $message.data('messageHelper').remove(function () {
					if (callback && {}.toString.call(callback) === '[object Function]') callback();
					if (adjust === true) adjustPage();
				});
			};

			var show = function ($message, callback, adjust) {
				return $message.data('messageHelper').show(function () {
					if (callback && {}.toString.call(callback) === '[object Function]') callback();
					if (adjust === true) adjustPage();
				});
			};

			var hide = function ($message, callback, adjust) {
				return $message.data('messageHelper').hide(function () {
					if (callback && {}.toString.call(callback) === '[object Function]') callback();
					if (adjust === true) adjustPage();
				});
			};

			var adjustPage = function () {
				var scroll = $window.scrollTop(),
					$site_content = $('#site_content'),
					height = $('#messages').outerHeight(true),
					currentHeight = parseInt($site_content.css('margin-top'));

				$site_content.css({ marginTop: height });
				$window.scrollTop(scroll + height - currentHeight);
			};

			// hooks
			initialize();

			// api
			this.data('messageHandler', {
				generate: generate,
				create: create,
				update: update,
				remove: remove,
				show: show,
				hide: hide
			});

			return this;
		},

		/**
		 *
		 * @param type
		 * @param args
		 * @returns {*}
		 */
		messageHelper: function (type, args) {
			var options,
				defaults = {
					speed: 150
				};

			var $message = $('<div class="message"><div class="wrap"></div></div>'),
				$this = this,
				message;

			// methods
			var test = function (message, type, args) {
				if (!validMessage(message))
					return false;

				if (!validType(type))
					return false;

				options = $.extend(defaults, args);

				return true;
			};

			var update = function (type, args) {
				var messageID = $this.attr('id'),
					message = $this.html();

				if (!test(message, type, args))
					return false;

				$this = $message.removeClass(types.join(' ')).addClass(type).attr('id', messageID);
				$this.find('.wrap').html(message);

				return $this;
			};

			var remove = function (callback) {
				return $this.hide(function () {
					$this.remove();
					callback();
				});
			};

			var show = function (callback) {
				return $this.slideDown(options.speed, callback);
			};

			var hide = function (callback) {
				return $this.slideUp(options.speed, callback);
			};

			// setup
			if (!update(type, args))
				return false;

			// api
			$this.data('messageHelper', {
				update: update,
				remove: remove,
				show: show,
				hide: hide
			});

			return $this;
		}
	});
})(jQuery, window, document);

/**
 * FORM HANDLER
 */
(function ($, window, document) {
	var $handler;

	$.extend($.fn, {
		formHandler: function (args) {
			//options
			args = $.extend({
				speed: 150
			}, args);

			$handler = this;


			// methods
			var initialize = function () {
				generate('form');
			};

			var generate = function (element) {
				var $forms = $(element),
					forms = {}, $form;

				$forms.each(function () {
					if ($form = $(this).formHelper(args))
						forms[$form.attr('id')] = $form;
				});

				return forms;
			};

			// hooks
			initialize();

			// api
			this.data('formHandler', {
				generate: generate
			});

			return this;
		},

		formHelper: function (args) {
			//options
			var options,
				defaults = {
					speed: 150
				};

			var $this = this,
				$messages = {},
				$repeaters = {};

			// methods
			var update = function (args) {
				options = $.extend(defaults, args);

				$messages = getMessages();
				$repeaters = getRepeaters();

				setClickers();
				setAjax();

				return $this;
			};

			var getMessages = function () {
				var $messages = $this.find('.message'),
					messages = {};

				$messages.each(function () {
					var $message = $(this),
						$input = $($message.parent().attr('data-for'));

					if ($messageHandler && ($message = $messageHandler.data('messageHandler').create($input, $message)))
						messages[$message.attr('id')] = $message;
				});

				return messages;
			};

			var getRepeaters = function () {
				var $repeaters = $this.find('.repeater'),
					repeaters = {};

				$repeaters.each(function () {
					var $repeater = $(this);

					if ($repeater = $repeater.formRepeater({}))
						repeaters[$repeater.attr('id')] = $repeater;
				});

				return repeaters;
			};

			var setClickers = function () {
				var check = function ($clicker) {
					var $input = $clicker.find('input');

					if ($input.prop('checked'))
						$clicker.addClass('check');
					else
						$clicker.removeClass('check');
				};

				$this.on('click', '.clicker', function (e) {
					var $clicker = $(this),
						$input = $clicker.find('input');

					$input.trigger('click');
					check($clicker);

					e.stopPropagation();
				});

				$this.find('.clicker').each(function () {
					check($(this));
				});
			};

			var $last = null;

			var setAjax = function () {
				if (!$this.hasClass('ajax'))
					return;

				$this.on('focus', ':input', function () {
					$last = $(this);
				});

				$this.submit(function (e) {
					var $this = $(this),
						$li = $this.closest('li'),
						liID = '#' + $li.attr('id');

					$.post($this.attr('action'), {
						name: $this.find('[name="name"]').val(),
						vote: $last.val()
					}, function (data) {
						var $data = $(data);

						$li.replaceWith($data.find(liID));

						$formHandler.data('formHandler').generate(liID + ' form');
					});

					e.preventDefault();
				});
			};

			// setup
			if (!update(args))
				return false;

			// api
			$this.data('formHelper', {
				update: update,
				messages: $messages
			});

			return $this;
		},

		formRepeater: function (args) {
			var $repeater = this,
				$repeats, repeatCount, $copy, newIndex = 0;

			var $tempControls = $('<div class="controls"></div>'),
				$tempControlsCreate = $('<div class="mod create">+</div>'),
				$tempControlsRemove = $('<div class="mod remove">-</div>'),
				$tempControlsMove = $('<div class="move"></div>');

			// methods
			var initialize = function () {
				updateRepeater($repeater);

				$repeats.each(function (index) {
					var $repeat = $(this),
						key = $repeat.find('label[for]').attr('for').split('__').shift().split('_').pop();

					if (key == repeatCount - 1)
						$copy = $repeat.clone();

					updateControls($repeat, index);

					newIndex++;
				});
			};

			var updateRepeater = function ($repeater) {
				$repeats = $repeater.find('.repeat');
				repeatCount = $repeats.length;
			};

			var updateControls = function ($repeat, index) {
				$repeat.find('.controls').remove();

				var $controls = $tempControls.clone();

				$controls.append($tempControlsRemove.clone());

				if (repeatCount >= 2)
					$controls.append($tempControlsMove.clone());

				if (index == repeatCount - 1)
					$controls.append($tempControlsCreate.clone());

				$repeat.append($controls);
			};

			var create = function ($repeater) {
				var cID = $copy.find('label[for]').attr('for').split('__').shift().split('_').pop();

				$repeater.append($($('<div>').append($copy).html().replace(new RegExp('_' + cID + '_', 'g'), '_' + newIndex + '_').replace(new RegExp('\\[' + cID + '\\]', 'g'), '[' + newIndex + ']')));

				updateRepeater($repeater);

				$repeats.each(function (index) {
					updateControls($(this), index);
				});

				newIndex++;
			};

			var remove = function ($repeater, $repeat) {
				if (repeatCount > 1)
					$repeat.remove();
				else
					$repeat.find('input, textarea').val('');

				updateRepeater($repeater);

				$repeats.each(function (index) {
					updateControls($(this), index);
				})
			};

			// setup
			initialize();

			$repeater.on('click', '.controls .remove', function () {
				remove($repeater, $(this).closest('.repeat'));
			});

			$repeater.on('click', '.controls .create', function () {
				create($repeater);
			});

			return $repeater;
		}
	});
})(jQuery, window, document);

/**
 * EASY FRAME
 */
(function ($, window, document) {
	var $handler;

	$.extend($.fn, {
		easyFrame: function (args) {
			var $easy = this,
				$frame = $easy.find('.frame'),
				$assets = $easy.find('.assets');

			var $currentView = null;

			// methods
			var initialize = function () {
				$frame.find('.vw').each(function () {
					console.log($(this));
					$(this).addClass('rlld').slideUp(0);
				});

				$assets.on('click', '.viewer', function (e) {
					var $this = $(this),
						$view = $frame.find($this.attr('data-target'));

					if ($this.hasClass('current')) {
						$currentView.fadeOut(150, function () {
							$currentView.fadeIn(150);
						});

//                        e.preventDefault();
						return;
					}

					if ($currentView)
						$currentView.slideUp(250, function () {
							$view.slideDown(250);
						});
					else
						$view.slideDown(250);

					$assets.find('.viewer').removeClass('current');
					$this.addClass('current');

					$currentView = $view;

//                    e.preventDefault();
				});

				$assets.find('.viewer').first().trigger('click');

				$frame.removeClass('hide');
			};

			// setup
			initialize();

			// api

			$easy.data('easyFrame', {

			});

			return $easy;
		}
	});
})(jQuery, window, document);

// initialization
var $messageHandler, $formHandler;

$(document).ready(function () {
	// global interfaces
	$messageHandler = $('#messages').messageHandler({});
	$formHandler = $('#forms').formHandler({});

	// everywhere
	dbjs.baselineFigure('.figure.scalar');
	dbjs.hoverable('.hoverable');
	dbjs.minimizable('.minimizable');
});
