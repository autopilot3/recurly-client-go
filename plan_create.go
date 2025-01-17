// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type PlanCreate struct {

	// Unique code to identify the plan. This is used in Hosted Payment Page URLs and in the invoice exports.
	Code *string `json:"code,omitempty"`

	// This name describes your plan and will appear on the Hosted Payment Page and the subscriber's invoice.
	Name *string `json:"name,omitempty"`

	// Optional description, not displayed.
	Description *string `json:"description,omitempty"`

	// Accounting code for invoice line items for the plan. If no value is provided, it defaults to plan's code.
	AccountingCode *string `json:"accounting_code,omitempty"`

	// Unit for the plan's billing interval.
	IntervalUnit *string `json:"interval_unit,omitempty"`

	// Length of the plan's billing interval in `interval_unit`.
	IntervalLength *int `json:"interval_length,omitempty"`

	// Units for the plan's trial period.
	TrialUnit *string `json:"trial_unit,omitempty"`

	// Length of plan's trial period in `trial_units`. `0` means `no trial`.
	TrialLength *int `json:"trial_length,omitempty"`

	// Allow free trial subscriptions to be created without billing info. Should not be used if billing info is needed for initial invoice due to existing uninvoiced charges or setup fee.
	TrialRequiresBillingInfo *bool `json:"trial_requires_billing_info,omitempty"`

	// Automatically terminate plans after a defined number of billing cycles.
	TotalBillingCycles *int `json:"total_billing_cycles,omitempty"`

	// Subscriptions will automatically inherit this value once they are active. If `auto_renew` is `true`, then a subscription will automatically renew its term at renewal. If `auto_renew` is `false`, then a subscription will expire at the end of its term. `auto_renew` can be overridden on the subscription record itself.
	AutoRenew *bool `json:"auto_renew,omitempty"`

	// A fixed pricing model has the same price for each billing period.
	// A ramp pricing model defines a set of Ramp Intervals, where a subscription changes price on
	// a specified cadence of billing periods. The price change could be an increase or decrease.
	PricingModel *string `json:"pricing_model,omitempty"`

	// Ramp Intervals
	RampIntervals []PlanRampIntervalCreate `json:"ramp_intervals,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomFieldCreate `json:"custom_fields,omitempty"`

	// Revenue schedule type
	RevenueScheduleType *string `json:"revenue_schedule_type,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	LiabilityGlAccountId *string `json:"liability_gl_account_id,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	RevenueGlAccountId *string `json:"revenue_gl_account_id,omitempty"`

	// The ID of a performance obligation. Performance obligations are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	PerformanceObligationId *string `json:"performance_obligation_id,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	SetupFeeLiabilityGlAccountId *string `json:"setup_fee_liability_gl_account_id,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	SetupFeeRevenueGlAccountId *string `json:"setup_fee_revenue_gl_account_id,omitempty"`

	// The ID of a performance obligation. Performance obligations are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	SetupFeePerformanceObligationId *string `json:"setup_fee_performance_obligation_id,omitempty"`

	// Setup fee revenue schedule type
	SetupFeeRevenueScheduleType *string `json:"setup_fee_revenue_schedule_type,omitempty"`

	// Accounting code for invoice line items for the plan's setup fee. If no value is provided, it defaults to plan's accounting code.
	SetupFeeAccountingCode *string `json:"setup_fee_accounting_code,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the plan is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraTransactionType *int `json:"avalara_transaction_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the plan is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraServiceType *int `json:"avalara_service_type,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's In-the-Box tax solution to determine taxation rules. You can pass in specific tax codes using any of these tax integrations. For Recurly's In-the-Box tax offering you can also choose to instead use simple values of `unknown`, `physical`, or `digital` tax codes.
	TaxCode *string `json:"tax_code,omitempty"`

	// `true` exempts tax on the plan, `false` applies tax on the plan.
	TaxExempt *bool `json:"tax_exempt,omitempty"`

	// Pricing
	Currencies []PlanPricingCreate `json:"currencies,omitempty"`

	// Hosted pages settings
	HostedPages *PlanHostedPagesCreate `json:"hosted_pages,omitempty"`

	// Add Ons
	AddOns []AddOnCreate `json:"add_ons,omitempty"`

	// Used to determine whether items can be assigned as add-ons to individual subscriptions.
	// If `true`, items can be assigned as add-ons to individual subscription add-ons.
	// If `false`, only plan add-ons can be used.
	AllowAnyItemOnSubscriptions *bool `json:"allow_any_item_on_subscriptions,omitempty"`

	// Unique ID to identify a dunning campaign. Used to specify if a non-default dunning campaign should be assigned to this plan. For sites without multiple dunning campaigns enabled, the default dunning campaign will always be used.
	DunningCampaignId *string `json:"dunning_campaign_id,omitempty"`
}
