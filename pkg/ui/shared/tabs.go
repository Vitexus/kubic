package shared

// TabItem is kind of identifier for tabs.
// It differs from `elements/tab` which is responsible for look and style, and which is data agnostic.
// TabItem, on the contrary, represents a concrete tab with a concrete title and related with a specific content.
type TabItem int

// TabItem tabs.
const (
	NamespacesTab TabItem = iota
	DeploymentsTab
	PodsTab
	AnyTab // used for elements that don't belong to any tab. As example, tabs themselves.
)

const (
	namespacesTabTitle  = "Namespaces"
	deploymentsTabTitle = "Deployments"
	podsTabTitle        = "Pods"
)

// String is a string representation of TabItems.
func (t TabItem) String() string {
	switch t {
	case NamespacesTab:
		return namespacesTabTitle
	case DeploymentsTab:
		return deploymentsTabTitle
	case PodsTab:
		return podsTabTitle
	default:
		return ""
	}
}

// GetTabItems returns the list of all available (visually) tabs.
func GetTabItems() []TabItem {
	return []TabItem{
		NamespacesTab,
		DeploymentsTab,
		PodsTab,
	}
}
