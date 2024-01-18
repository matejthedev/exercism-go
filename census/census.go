// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && r.Address["street"] != ""
}

func (r *Resident) Delete() {
	r.Address = nil
	r.Age = 0
	r.Name = ""
}

func Count(residents []*Resident) int {
	count := 0
	for _, v := range residents {
		if v.HasRequiredInfo() {
			count++
		}
	}
	return count
}
