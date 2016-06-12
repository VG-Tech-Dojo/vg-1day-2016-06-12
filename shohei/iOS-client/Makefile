XCODEBUILD?=$(shell which xcodebuild)
SCHEME?="OneDayInternshipApp"
DESTINATION?='platform=iOS Simulator,OS=9.3,name=iPhone 6s'
CONFIGURATION?=Test
PROJ?="OneDayInternshipApp.xcodeproj"

all:

help:
	cat Makefile

test:
	$(XCODEBUILD) \
		-project $(PROJ) \
		-scheme $(SCHEME) \
		-destination $(DESTINATION) \
		-configuration $(CONFIGURATION) \
		test
