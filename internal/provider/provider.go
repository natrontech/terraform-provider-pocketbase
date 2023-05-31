package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/natrontech/pocketbase-client-go/pkg/client"
)

// Ensure PocketbaseProvider satisfies various provider interfaces.
var _ provider.Provider = &PocketbaseProvider{}

// PocketbaseProvider defines the provider implementation.
type PocketbaseProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// PocketbaseProviderModel describes the provider data model.
type pocketbaseProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Identity types.String `tfsdk:"identity"`
	Password types.String `tfsdk:"password"`
}

func (p *PocketbaseProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pocketbase"
	resp.Version = p.version
}

func (p *PocketbaseProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Pocketbase provider",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The endpoint to use for the Pocketbase API (e.g. http://127.0.0.1:8090)",
				Required:            true,
			},
			"identity": schema.StringAttribute{
				MarkdownDescription: "The identity to use for the Pocketbase API (e.g. admin@natron.io)",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The password to use for the Pocketbase API",
				Required:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *PocketbaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Pocketbase provider")

	// Retrieve provider data from configuration
	var config pocketbaseProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown Pocketbase endpoint",
			"The provider cannot create the Pocketbase client because the endpoint is unknown",
		)
	}

	if config.Identity.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("identity"),
			"Unknown Pocketbase identity",
			"The provider cannot create the Pocketbase client because the identity is unknown",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Pocketbase password",
			"The provider cannot create the Pocketbase client because the password is unknown",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := os.Getenv("POCKETBASE_ENDPOINT")
	identity := os.Getenv("POCKETBASE_IDENTITY")
	password := os.Getenv("POCKETBASE_PASSWORD")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	if !config.Identity.IsNull() {
		identity = config.Identity.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown Pocketbase endpoint",
			"The provider cannot create the Pocketbase client because the endpoint is unknown",
		)
	}

	if identity == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("identity"),
			"Unknown Pocketbase identity",
			"The provider cannot create the Pocketbase client because the identity is unknown",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Pocketbase password",
			"The provider cannot create the Pocketbase client because the password is unknown",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "pocketbase_endpoint", endpoint)
	ctx = tflog.SetField(ctx, "pocketbase_identity", identity)
	ctx = tflog.SetField(ctx, "pocketbase_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "pocketbase_password")

	tflog.Debug(ctx, "Creating Pocketbase client")

	// Create the Pocketbase client
	client, err := client.NewClient(&endpoint, &identity, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Pocketbase client",
			"Unable to create Pocketbase client: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Pocketbase provider", map[string]any{"success": true})
}

func (p *PocketbaseProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// NewExampleResource,
	}
}

func (p *PocketbaseProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// NewExampleDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PocketbaseProvider{
			version: version,
		}
	}
}
