package libmedialog

import (
	"github.com/google/uuid"
	"time"
)

type Accession struct {
	ID               int       `json:"id"`
	AccessionNum     string    `json:"accession_num"`
	AccessionNote    string    `json:"accession_note"`
	AccessionState   string    `json:"accession_state"`
	CollectionID     int       `json:"collection_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedBy        int       `json:"created_by"`
	ModifiedBy       int       `json:"modified_at"`
	ASpaceRepository int       `json:"aspace_repository"`
	AspaceID         int       `json:"aspace_id"`
}

type Collection struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	CollectionCode string `json:"collection_code"`
	PartnerCode    string `json:"partner_code"`
}

type MlogEntry struct {
	ID                    uuid.UUID `json:"id"`
	AccessionNum          string    `json:"accession_num"`
	MediaID               int       `json:"media_id"`
	MediaType             string    `json:"media_type"`
	Manufacturer          string    `json:"manufacturer"`
	ManufacturerSerial    string    `json:"manufacturer_serial"`
	LabelText             string    `json:"label_text"`
	MediaNote             string    `json:"media_note"`
	PhotoURL              string    `json:"photo_url"`
	ImageFilename         string    `json:"image_filename"`
	CPUInterface          string    `json:"interface"`
	ImagingSoftware       string    `json:"imaging_software"`
	HDDInterface          string    `json:"hdd_interface"`
	ImagingSuccess        string    `json:"imaging_success"`
	ImagedBy              string    `json:"imaged_by"`
	InterpretationSuccess string    `json:"interpretation_success"`
	ImagingNote           string    `json:"imaging_note"`
	ImageFormat           string    `json:"image_format"`
	EncodingScheme        string    `json:"encoding_scheme"`
	PartitionTableFormat  string    `json:"partition_table_format"`
	NumberOfPartitions    int       `json:"number_of_partitions"`
	Filesystem            string    `json:"filesystem"`
	HasDFXML              bool      `json:"has_dfxml"`
	HasFTKCSV             bool      `json:"has_ftkcsv"`
	ImageSizeBytes        int64     `json:"image_size_bytes"`
	MD5Checksum           string    `json:"md5_checksum"`
	SHA1Checksum          string    `json:"sha_1_checksum"`
	DateImaged            time.Time `json:"date_imaged"`
	DateFTKLoaded         time.Time `json:"date_ftk_loaded"`
	DateMetadataExtracted time.Time `json:"date_metadata_extracted"`
	DatePhotographed      time.Time `json:"date_photographed"`
	DateQC                time.Time `json:"date_qc"`
	DatePackaged          time.Time `json:"date_packaged"`
	DateTransferred       time.Time `json:"date_transferred"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAT             time.Time `json:"updated_at"`
	NumberImageSegments   int       `json:"number_of_image_segments"`
	RefID                 string    `json:"ref_id"`
	HasMactimeCSV         bool      `json:"has_mactime_csv"`
	BoxNumber             int       `json:"box_number"`
	StockSize             string    `json:"stock_size"`
	SIPID                 int       `json:"sip_id"`
	Original_ID           string    `json:"original_id"`
	DispositionNote       string    `json:"disposition_note"`
	StockUnit             string    `json:"stock_unit"`
	StockSizeNum          float32   `json:"stock_size_num"`
	CreatedBy             int       `json:"created_by"`
	ModifiedBy            int       `json:"modified_by"`
	CollectionID          int       `json:"collection_id"`
	AccessionID           int       `json:"accession_id"`
	RepositoryID          int       `json:"repository_id"`
	IsTransferred         bool      `json:"is_transferred"`
	IsRefreshed           bool      `json:"is_refreshed"`
	SessionCount          int       `json:"session_count"`
	ContentType           string    `json:"content_type"`
	Structure             string    `json:"structure"`
	FileSystems           string    `json:"file_systems"`
}

type User struct {
	ID                  int       `json:"id"`
	Email               string    `json:"email"`
	EncryptedPassword   string    `json:"encrypted_password"`
	ResetPasswordToken  string    `json:"reset_password_token"`
	ResetPasswordSentAt time.Time `json:"reset_password_sent_at"`
	RememberCreatedAt   time.Time `json:"remember_created_at"`
	SignInCount         string    `json:"sign_in_count"`
	LastSignInAt        time.Time `json:"last_sign_in_at"`
	CurrentSignInIP     string    `json:"current_sign_in_ip"`
	LastSignInIP        string    `json:"last_sign_in_ip"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
	Admin               bool      `json:"admin"`
	IsActive            bool      `json:"is_active"`
}
