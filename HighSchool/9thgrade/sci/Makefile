all: render
    
render: report.1
	groff -ms -Tutf8 $< | less

print: report.1
	groff -pspic -ms $< > report.ps
	lp report.ps
	rm report.ps
edit:
	/usr/bin/vim report.1

ps:
	groff -ms -Tps -fA report.1 > psreport.ps

evince:  ps
	evince psreport.ps

pdf: 
	groff -ms -Tpdf -fC report.1 > psreport.pdf
