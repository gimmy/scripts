(setq make-backup-files nil)
(tool-bar-mode  -1)
(menu-bar-mode -1)
(scroll-bar-mode -1)

(global-visual-line-mode)
(setq auto-mode-alist (cons '( "\\.m" . octave-mode) auto-mode-alist ) )
(setq auto-mode-alist (cons '( "\\.less" . css-mode) auto-mode-alist ) )

(autoload 'php-mode "php-mode" "Major mode for editing php code." t)
(add-to-list 'auto-mode-alist '("\\.php$" . php-mode))
(add-to-list 'auto-mode-alist '("\\.inc$" . php-mode))

(setq auto-mode-alist (cons '("\\.ml[iylp]?\\'" . tuareg-mode) auto-mode-alist))
  (autoload 'tuareg-mode "tuareg" "Major mode for editing Caml code" t)
    (autoload 'ocamldebug "ocamldebug" "Run the Caml debugger" t)
    
;; (add-to-list 'load-path "/usr/share/doc/git/contrib/emacs")
;; (require 'git)
;; (require 'git-blame)

(load "auctex.el" nil t t)
(load "preview-latex.el" nil t t)

;; YASnippet
(add-to-list 'load-path "/usr/share/emacs/site-lisp/yas")
(require 'yasnippet) ;; not yasnippet-bundle
;; (yas/initialize)
;; (yas/global-mode 1)

;; PKGBUILD Mode
(autoload 'pkgbuild-mode "pkgbuild-mode.el" "PKGBUILD mode." t)
(setq auto-mode-alist (append '(("/PKGBUILD$" . pkgbuild-mode)) auto-mode-alist))

;; Vala Mode
(autoload 'vala-mode "vala-mode" "Major mode for editing Vala code." t)
(add-to-list 'auto-mode-alist '("\.vala$" . vala-mode))
(add-to-list 'auto-mode-alist '("\.vapi$" . vala-mode))
(add-to-list 'file-coding-system-alist '("\.vala$" . utf-8))
(add-to-list 'file-coding-system-alist '("\.vapi$" . utf-8))


; Macro to load somethings
(defun load-book ()
  (interactive)
  (desktop-change-dir "~/git/apnea")
  (desktop-save-mode 1)
  )
(defun load-DBLP ()
  (interactive)
  (desktop-change-dir "~/git/DBLP")
  (desktop-save-mode 1)
  )
(defun load-tw ()
  (interactive)
  (desktop-change-dir "~/git/tweetcionary")
  (desktop-save-mode 1)
  )
(defun  bhalostrap ()
  (interactive)
  (desktop-change-dir "~/Web/www.bhalobasa.it/wordpress/wp-content/themes/the-bootstrap")
  (desktop-save-mode 1)
  )

(add-to-list 'load-path "~/.emacs.d")

;; Color theme
(require 'color-theme)
;;(setq color-theme-load-all-themes nil)
;;(require 'color-theme-tangotango)


;; select theme - first list element is for windowing system, second is for console/terminal
;; Source : http://www.emacswiki.org/emacs/ColorTheme#toc9
;; (if window-system
;;     (color-theme-initialize)
;;     (color-theme-comidia))

;; (setq color-theme-choices 
;;      '(color-theme-comidia color-theme-comidia))

;;(color-theme-tangotango)
(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(custom-enabled-themes (quote (adwaita)))
 '(inhibit-startup-screen t))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )
