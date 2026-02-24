package msi

type Category int

const (
	CategoryText = iota
	CategoryUpperCase
	CategoryLowerCase
	CategoryInteger
	CategoryDoubleInteger
	CategoryTimeDate
	CategoryIdentifier
	CategoryProperty
	CategoryFilename
	CategoryWildCardFilename
	CategoryPath
	CategoryPaths
	CategoryAnyPath
	CategoryDefaultDir
	CategoryRegPath
	CategoryFormatted
	CategoryFormattedSddlText
	CategoryTemplate
	CategoryCondition
	CategoryGuid
	CategoryVersion
	CategoryLanguage
	CategoryBinary
	CategoryCustomSource
	CategoryCabinet
	CategoryShortcut
)

var AllCategories = []Category{
	CategoryText, CategoryUpperCase, CategoryLowerCase, CategoryInteger,
	CategoryDoubleInteger, CategoryTimeDate, CategoryIdentifier, CategoryProperty,
	CategoryFilename, CategoryWildCardFilename, CategoryPath, CategoryPaths,
	CategoryAnyPath, CategoryDefaultDir, CategoryRegPath, CategoryFormatted,
	CategoryFormattedSddlText, CategoryTemplate, CategoryCondition, CategoryGuid,
	CategoryVersion, CategoryLanguage, CategoryBinary, CategoryCustomSource,
	CategoryCabinet, CategoryShortcut,
}

func (c Category) String() string {
	switch c {
	case CategoryAnyPath:
		return "AnyPath"
	case CategoryBinary:
		return "Binary"
	case CategoryCabinet:
		return "Cabinet"
	case CategoryCondition:
		return "Condition"
	case CategoryCustomSource:
		return "CustomSource"
	case CategoryDefaultDir:
		return "DefaultDir"
	case CategoryDoubleInteger:
		return "DoubleInteger"
	case CategoryFilename:
		return "Filename"
	case CategoryFormatted:
		return "Formatted"
	case CategoryFormattedSddlText:
		return "FormattedSddlText"
	case CategoryGuid:
		return "GUID"
	case CategoryIdentifier:
		return "Identifier"
	case CategoryInteger:
		return "Integer"
	case CategoryLanguage:
		return "Language"
	case CategoryLowerCase:
		return "LowerCase"
	case CategoryPath:
		return "Path"
	case CategoryPaths:
		return "Paths"
	case CategoryProperty:
		return "Property"
	case CategoryRegPath:
		return "RegPath"
	case CategoryShortcut:
		return "Shortcut"
	case CategoryTemplate:
		return "Template"
	case CategoryText:
		return "Text"
	case CategoryTimeDate:
		return "TimeDate"
	case CategoryUpperCase:
		return "UpperCase"
	case CategoryVersion:
		return "Version"
	case CategoryWildCardFilename:
		return "WildCardFilename"
	default:
		return ""
	}
}

func CategoryFromString(str string) Category {
	switch str {
	case "AnyPath":
		return CategoryAnyPath
	case "Binary":
		return CategoryBinary
	case "Cabinet":
		return CategoryCabinet
	case "Condition":
		return CategoryCondition
	case "CustomSource":
		return CategoryCustomSource
	case "DefaultDir":
		return CategoryDefaultDir
	case "DoubleInteger":
		return CategoryDoubleInteger
	case "Filename":
		return CategoryFilename
	case "Formatted":
		return CategoryFormatted
	case "FormattedSddlText":
		return CategoryFormattedSddlText
	case "GUID", "Guid":
		return CategoryGuid
	case "Identifier":
		return CategoryIdentifier
	case "Integer":
		return CategoryInteger
	case "Language":
		return CategoryLanguage
	case "LowerCase":
		return CategoryLowerCase
	case "Path":
		return CategoryPath
	case "Paths":
		return CategoryPaths
	case "Property":
		return CategoryProperty
	case "RegPath":
		return CategoryRegPath
	case "Shortcut":
		return CategoryShortcut
	case "Template":
		return CategoryTemplate
	case "Text":
		return CategoryText
	case "TimeDate":
		return CategoryTimeDate
	case "UpperCase":
		return CategoryUpperCase
	case "Version":
		return CategoryVersion
	case "WildCardFilename":
		return CategoryWildCardFilename
	default:
		return -1
	}
}
