package generator

// Options represents the options for the generator.
type Options struct {
	AppName    string
	PageName   string
	PageNameUp string

	Directory string
}

// Option manipulates the Options passed.
type Option func(o *Options)

// AppName sets the app name.
func AppName(s string) Option {
	return func(o *Options) {
		o.AppName = s
	}
}

// PageName sets the page name.
func PageName(s string) Option {
	return func(o *Options) {
		o.PageName = s
	}
}

// PageNameUp sets the page name.
func PageNameUp(s string) Option {
	return func(o *Options) {
		o.PageNameUp = s
	}
}

// Directory sets the directory in which files are generated.
func Directory(d string) Option {
	return func(o *Options) {
		o.Directory = d
	}
}
