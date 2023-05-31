package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/natrontech/pocketbase-client-go/pkg/client"
)

var (
	_ resource.Resource                = &collectionResource{}
	_ resource.ResourceWithConfigure   = &collectionResource{}
	_ resource.ResourceWithImportState = &collectionResource{}
)

// NewCollectionResource is a helper function to simplify the provider implementation.
func NewCollectionResource() resource.Resource {
	return &collectionResource{}
}

type collectionResource struct {
	client *client.Client
}

type collectionResourceModel struct {
	Id         types.String                    `tfsdk:"id"`
	Name       types.String                    `tfsdk:"name"`
	Type       types.String                    `tfsdk:"type"`
	Schema     []collectionResourceSchemaModel `tfsdk:"schema"`
	System     types.Bool                      `tfsdk:"system"`
	ListRule   types.Bool                      `tfsdk:"list_rules"`
	ViewRule   types.Bool                      `tfsdk:"view_rules"`
	CreateRule types.Bool                      `tfsdk:"create_rules"`
	UpdateRule types.Bool                      `tfsdk:"update_rules"`
	DeleteRule types.Bool                      `tfsdk:"delete_rules"`
	Options    []types.List                    `tfsdk:"options"`
	Indexes    []types.String                  `tfsdk:"indexes"`
	Created    types.String                    `tfsdk:"created"`
	Updated    types.String                    `tfsdk:"updated"`
}

type collectionResourceSchemaModel struct {
	System   types.Bool   `tfsdk:"system"`
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	Required types.Bool   `tfsdk:"required"`
	Unique   types.Bool   `tfsdk:"unique"`
	Options  []types.List `tfsdk:"options"`
}

// Metadata returns the resource type name.
func (r *collectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_collection"
}

// Schema defines the schem for the resource.
func (r *collectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Pocketbase collection",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *collectionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got %T", req.ProviderData),
		)
	}

	r.client = client
}

// Create a new resource.
func (e *collectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data collectionResourceModel

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create resource using 3rd party API.

	data.Id = types.StringValue("example-id")

	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Read resource information.
func (e *collectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data collectionResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read resource using 3rd party API.

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Update a resource.
func (e *collectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data collectionResourceModel

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update resource using 3rd party API.

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Delete a resource.
func (e *collectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data collectionResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete resource using 3rd party API.
}

// Import a resource.
func (r *collectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
