.PHONY: styles
styles:
	npx tailwindcss -i assets/tailwind.css -o assets/dist/styles.css

.PHONY: templates
templates: styles
	templ generate