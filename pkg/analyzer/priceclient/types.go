package apis

type RegionTypeKey struct {
	Region       string
	InstanceType string
}

type RegionalEC2Price struct {
	InstanceTypeEC2Price map[string]*InstanceTypePrice `json:"instanceTypeEC2Price"`
}

type InstanceTypePrice struct {
	Arch                 string   `json:"arch"`
	VCPU                 float64  `json:"vcpu"`
	Memory               float64  `json:"memory"`
	GPU                  float64  `json:"gpu"`
	Zones                []string `json:"zones"`
	OnDemandPricePerHour float64  `json:"onDemandPricePerHour"`
	// AWSEC2Billing represents the cost of saving plan billing
	// key is {savings plan type}/{term length}/{payment option}
	AWSEC2Billing map[string]AWSEC2Billing `json:"awsEC2Billing"`
	// SpotPricePerHour represents the smallest spot price per hour in different zones
	SpotPricePerHour map[string]float64 `json:"spotPricePerHour"`
}

type AWSEC2Billing struct {
	Rate float64 `json:"rate"`
}

type AWSEC2SPPaymentOption string

const (
	AWSEC2SPPaymentOptionAllUpfront     AWSEC2SPPaymentOption = "all"
	AWSEC2SPPaymentOptionPartialUpfront AWSEC2SPPaymentOption = "partial"
	AWSEC2SPPaymentOptionNoUpfront      AWSEC2SPPaymentOption = "no"
)

func (r *RegionalEC2Price) DeepCopy() *RegionalEC2Price {
	d := &RegionalEC2Price{
		InstanceTypeEC2Price: make(map[string]*InstanceTypePrice),
	}
	for k, v := range r.InstanceTypeEC2Price {
		vCopy := v.DeepCopy()
		d.InstanceTypeEC2Price[k] = vCopy
	}
	return d
}

func (i *InstanceTypePrice) DeepCopy() *InstanceTypePrice {
	d := &InstanceTypePrice{
		Arch:                 i.Arch,
		VCPU:                 i.VCPU,
		Memory:               i.Memory,
		GPU:                  i.GPU,
		Zones:                make([]string, len(i.Zones)),
		OnDemandPricePerHour: i.OnDemandPricePerHour,
		AWSEC2Billing:        make(map[string]AWSEC2Billing),
		SpotPricePerHour:     make(map[string]float64),
	}
	copy(d.Zones, i.Zones)
	for k, v := range i.AWSEC2Billing {
		d.AWSEC2Billing[k] = v
	}
	for k, v := range i.SpotPricePerHour {
		d.SpotPricePerHour[k] = v
	}
	return d
}
