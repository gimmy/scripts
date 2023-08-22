;; General settings
(setq make-backup-files nil)
(tool-bar-mode  -1)
(menu-bar-mode -1)
(scroll-bar-mode -1)

(global-visual-line-mode)
(setq auto-mode-alist (cons '( "\\.m" . octave-mode) auto-mode-alist ) )
(setq auto-mode-alist (cons '( "\\.less" . css-mode) auto-mode-alist ) )

;; Packages: MELPA stable
(require 'package)
(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/") t)
;; Comment/uncomment this line to enable MELPA Stable if desired.  See `package-archive-priorities`
;; and `package-pinned-packages`. Most users will not need or want to do this.
;;(add-to-list 'package-archives '("melpa-stable" . "https://stable.melpa.org/packages/") t)
(package-initialize)

;; Load Tango color theme theme
(load-theme 'tangotango t)

;; PHP
(autoload 'php-mode "php-mode" "Major mode for editing php code." t)
(add-to-list 'auto-mode-alist '("\\.php$" . php-mode))
(add-to-list 'auto-mode-alist '("\\.inc$" . php-mode))

;; OCaml
(setq auto-mode-alist (cons '("\\.ml[iylp]?\\'" . tuareg-mode) auto-mode-alist))
  (autoload 'tuareg-mode "tuareg" "Major mode for editing Caml code" t)
    (autoload 'ocamldebug "ocamldebug" "Run the Caml debugger" t)
    
;; YASnippet
(add-to-list 'load-path "~/.emacs.d/elpa/yasnippet-0.11.0/")
;; (add-to-list 'yas-snippet-dirs "~/.emacs.d/elpa/yasnippet-0.11.0/yasnippet-snippets/")
(require 'yasnippet)
(yas-global-mode t)

;; LaTeX
(add-to-list 'load-path "~/.emacs.d/magic-latex-buffer/")
(require 'tex-mode)
(require 'magic-latex-buffer)
(add-hook 'latex-mode-hook 'magic-latex-buffer 'flyspell-mode)
(add-hook 'LaTeX-mode-hook 'magic-latex-buffer 'flyspell-mode)
;; (load "auctex.el" nil t t)
;; (load "preview-latex.el" nil t t)
(add-hook 'LaTeX-mode-hook
      '(lambda ()
         (TeX-add-symbols '("eqref" TeX-arg-ref (ignore)))))

;; Vala Mode
(autoload 'vala-mode "vala-mode" "Major mode for editing Vala code." t)
(add-to-list 'auto-mode-alist '("\.vala$" . vala-mode))
(add-to-list 'auto-mode-alist '("\.vapi$" . vala-mode))
(add-to-list 'file-coding-system-alist '("\.vala$" . utf-8))
(add-to-list 'file-coding-system-alist '("\.vapi$" . utf-8))

;; Markdown-mode
(autoload 'markdown-mode "markdown-mode" "Major mode for editing markdown." t)
(add-to-list 'auto-mode-alist '("\\.md$" . markdown-mode))

; Macro to load somethings
(defun load-diary ()
  (interactive)
  (desktop-change-dir "~/git/diary")
  (desktop-save-mode 1)
  )
(defun load-Web ()
  (interactive)
  (desktop-change-dir "~/Web/public_hugo")
  (desktop-save-mode 1)
  )

;; hook on after-make-frame-functions
(add-hook 'after-make-frame-functions 'test-win-sys)

(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(inhibit-startup-screen t)
 '(package-selected-packages
   '(dictionary org-roam-ui org-roam bibretrieve markdown-mode magit yasnippet auctex)))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )
