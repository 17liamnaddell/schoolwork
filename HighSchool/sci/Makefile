all: render
    
render: report.1
	groff -ms -Tutf8 $< | less

print: report.1
	groff -pspic -ms $< > report.ps
	lp report.ps
	rm report.ps
edit:
	/usr/bin/vim report.1

evince: report.1
	groff -ms -Tps -fA $< > psreport.ps
	evince psreport.ps
