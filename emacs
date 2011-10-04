(setq make-backup-file nil)
;; (tool-bar-mode nil)
;; (scroll-bar-mode nil)

(global-visual-line-mode)
(setq auto-mode-alist (cons '( "\\.m" . octave-mode) auto-mode-alist ) )

;; (require 'yasnippet-bundle)
;; (add-to-list 'load-path "/usr/share/doc/git/contrib/emacs")
;; (require 'git)
;; (require 'git-blame)

; Macro to load clinica site
(defun load-siteclinica ()
  (interactive)
  (desktop-change-dir "~/public_html/siteclinica")
  (desktop-save-mode 1)
  )

(add-to-list 'load-path "~/.emacs.d")

;; Color theme
(require 'color-theme)
(setq color-theme-load-all-themes nil)
(require 'color-theme-tangotango)

;; select theme - first list element is for windowing system, second is for console/terminal
;; Source : http://www.emacswiki.org/emacs/ColorTheme#toc9
(setq color-theme-choices 
      '(color-theme-tangotango color-theme-tangotango))

(color-theme-tangotango)
