package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/natrontech/pocketbase-client-go/pkg/client"
	"github.com/natrontech/pocketbase-client-go/pkg/models"
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
	models.Collection
}

// Metadata returns the resource type name.
func (r *collectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_collection"
}

// Schema defines the schem for the resource
func (r *collectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Pocketbase collection",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the collection",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created": schema.StringAttribute{
				Description: "The creation date of the collection",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "The last update date of the collection",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the collection",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the collection",
				Required:    true,
			},
			"schema": schema.ListNestedAttribute{
				Description: "The schema of the collection",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"system": schema.BoolAttribute{
							Description: "Whether the field is a system field",
							Required:    true,
						},
						"id": schema.StringAttribute{
							Description: "The ID of the field",
							Required:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the field",
							Required:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of the field",
							Required:    true,
						},
						"required": schema.BoolAttribute{
							Description: "Whether the field is required",
							Required:    true,
						},
						"unique": schema.BoolAttribute{
							Description: "Whether the field is unique",
							Required:    true,
						},
						"options": schema.ListAttribute{
							Description: "The options of the field",
							Required:    true,
						},
					},
				},
			},
			"list_rule": schema.StringAttribute{
				Description: "The list rule of the collection",
				Required:    true,
			},
			"view_rule": schema.StringAttribute{
				Description: "The view rule of the collection",
				Required:    true,
			},
			"create_rule": schema.StringAttribute{
				Description: "The create rule of the collection",
				Required:    true,
			},
			"update_rule": schema.StringAttribute{
				Description: "The update rule of the collection",
				Required:    true,
			},
			"delete_rule": schema.StringAttribute{
				Description: "The delete rule of the collection",
				Required:    true,
			},
			"options": schema.ListAttribute{
				Description: "The options of the collection",
				Required:    true,
			},
			"indexes": schema.ListAttribute{
				Description: "The indexes of the collection",
				Required:    true,
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
func (r *collectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan collectionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var collectionRequest models.CollectionCreateRequest
	collectionRequest.Id = &plan.Id
	collectionRequest.Name = plan.Name
	collectionRequest.Type = plan.Type
	collectionRequest.Schema = plan.Schema
	collectionRequest.ListRule = &plan.ListRule
	collectionRequest.ViewRule = &plan.ViewRule
	collectionRequest.CreateRule = &plan.CreateRule
	collectionRequest.UpdateRule = &plan.UpdateRule
	collectionRequest.DeleteRule = &plan.DeleteRule
	collectionRequest.Options = &plan.Options
	collectionRequest.Indexes = &plan.Indexes

	// Create the resource
	collection, err := r.client.CreateCollection(&collectionRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating collection",
			fmt.Sprintf("Error creating collection: %s", err),
		)
		return
	}

	// Set the ID of the created resource
	diags = resp.State.Set(ctx, collectionResourceModel{
		// TODO: Test if it works
		Collection: *collection,
	})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *collectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// TODO: Implement
}

// Update a resource.
func (r *collectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// TODO: Implement
}

// Delete a resource.
func (r *collectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// TODO: Implement
}

// Import a resource.
func (r *collectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
