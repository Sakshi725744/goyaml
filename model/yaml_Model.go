package model

type SpdxYaml struct {
	SpdxId string `yaml:"SPDXID"`

	SpdxVersion string `yaml:"spdxVersion"`

	CreationInfo struct {
		Comment            string   `yaml:"comment"`
		Created            string   `yaml:"created"`
		Creators           []string `yaml:"creators"`
		LicenseListVersion string   `yaml:"licenseListVersion"`
	} `yaml:"creationInfo"`

	Name string `yaml:"name"`

	DataLicense string `yaml:"dataLicense"`

	Comment string `yaml:"comment"`

	ExternalDocumentRefs []struct {
		ExternalDocumentId string `yaml:"externalDocumentId"`
		Checksum           struct {
			Algorithm     string `yaml:"algorithm"`
			ChecksumValue string `yaml:"checksumValue"`
		} `yaml:"checksum"`
		SpdxDocument string `yaml:"spdxDocument"`
	} `yaml:"externalDocumentRefs"`

	HasExtractedLicensingInfos []struct {
		ExtractedTexts string   `yaml:"extractedTexts"`
		LicenseId      string   `yaml:"licenseId"`
		Comment        string   `yaml:"comment,omitempty"`
		Name           string   `yaml:"name,omitempty"`
		SeeAlsos       []string `yaml:"seeAlsos,omitempty"`
	} `yaml:"hasExtractedLicensingInfos"`

	Annotations []struct {
		AnnotationDate string `yaml:"annotationDate"`
		AnnotationType string `yaml:"annotationType"`
		Annotator      string `yaml:"annotator"`
		Comment        string `yaml:"comment"`
	} `yaml:"annotations"`

	DocumentNamespace string `yaml:"documentNamespace"`

	DocumentDescribes []string `yaml:"documentDescribes"`

	Packages []struct {
		SpdxId      string `yaml:"SPDXID"`
		Annotations []struct {
			AnnotationDate string `yaml:"annotationDate"`
			AnnotationType string `yaml:"annotationType"`
			Annotator      string `yaml:"annotator"`
			Comment        string `yaml:"comment"`
		} `yaml:"annotations,omitempty"`
		Checksums []struct {
			Algorithm     string `yaml:"algorithm"`
			ChecksumValue string `yaml:"checksumValue"`
		} `yaml:"checksums,omitempty"`
		CopyrightText    string `yaml:"copyrightText,omitempty"`
		Description      string `yaml:"description,omitempty"`
		DownloadLocation string `yaml:"downloadLocation,omitempty"`
		ExternalRefs     []struct {
			Comment           string `yaml:"comment,omitempty"`
			ReferenceCategory string `yaml:"referenceCategory"`
			ReferenceLocator  string `yaml:"referenceLocator"`
			ReferenceType     string `yaml:"referenceType"`
		} `yaml:"externalRefs,omitempty"`
		FilesAnalyzed           bool     `yaml:"filesAnalyzed,omitempty"`
		HasFiles                []string `yaml:"hasFiles,omitempty"`
		Homepage                string   `yaml:"homepage,omitempty"`
		LicenseComments         string   `yaml:"licenseComments,omitempty"`
		LicenseConcluded        string   `yaml:"licenseConcluded,omitempty"`
		LicenseDeclared         string   `yaml:"licenseDeclared,omitempty"`
		LicenseInfoInFiles      []string `yaml:"licenseInfoInFiles,omitempty"`
		Name                    string   `yaml:"name,omitempty"`
		Originator              string   `yaml:"originator,omitempty"`
		PackageFileName         string   `yaml:"packageFileName"`
		PackageVerificationCode struct {
			PackageVerificationCodeExcludedFiles []string `yaml:"packageVerificationCodeExcludedFiles"`
			PackageVerificationCodeValue         string   `yaml:"packageVerificationCodeValue"`
		} `yaml:"packageVerificationCode,omitempty"`
		SourceInfo  string `yaml:"sourceInfo,omitempty"`
		Summary     string `yaml:"summary,omitempty"`
		Supplier    string `yaml:"supplier,omitempty"`
		VersionInfo string `yaml:"versionInfo,omitempty"`
	}

	Files []struct {
		SpdxId      string `yaml:"SPDXID"`
		Annotations []struct {
			AnnotationDate string `yaml:"annotationDate"`
			AnnotationType string `yaml:"annotationType"`
			Annotator      string `yaml:"annotator"`
			Comment        string `yaml:"comment"`
		} `yaml:"annotations,omitempty"`
		Checksums []struct {
			Algorithm     string `yaml:"algorithm"`
			ChecksumValue string `yaml:"checksumValue"`
		} `yaml:"checksums"`
		Comment            string   `yaml:"comment,omitempty"`
		CopyrightText      string   `yaml:"copyrightText "`
		FileContributors   []string `yaml:"fileContributors"`
		FileDependencies   []string `yaml:"fileDependencies"`
		FileName           string   `yaml:"fileName"`
		FileTypes          []string `yaml:"fileTypes"`
		LicenseComments    string   `yaml:"licenseComments,omitempty"`
		LicenseConcluded   string   `yaml:"licenseConcluded"`
		LicenseInfoInFiles []string `yaml:"licenseInfoInFiles"`
		NoticeText         string   `yaml:"noticeText,omitempty"`
	}

	Snippets []struct {
		SpdxId             string   `yaml:"SPDXID"`
		Comment            string   `yaml:"comment,omitempty"`
		CopyrightText      string   `yaml:"copyrightText "`
		LicenseComments    string   `yaml:"licenseComments,omitempty"`
		LicenseConcluded   string   `yaml:"licenseConcluded"`
		LicenseInfoInFiles []string `yaml:"licenseInfoInFiles"`
		Name               string   `yaml:"name"`
		Ranges             []struct {
			EndPointer struct {
				Offset    int    `yaml:"offset"`
				Reference string `yaml:"reference"`
			} `yaml:"endPointer"`
			StartPointer struct {
				Offset    int    `yaml:"offset"`
				Reference string `yaml:"reference"`
			} `yaml:"name"`
		} `yaml:"startPointer"`
		SnippetFromFile string `yaml:"snippetFromFile"`
	}

	Relationships []struct {
		SpdxElementId      string `yaml:"spdxElementId"`
		RelatedSpdxElement string `yaml:"relatedSpdxElement"`
		RelationshipType   string `yaml:"relationshipType"`
	} `yaml:"relationships"`
}
