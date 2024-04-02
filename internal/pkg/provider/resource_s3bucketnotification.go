package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceS3BucketNotification() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceS3BucketNotificationCreate,
		ReadContext:   resourceS3BucketNotificationRead,
		UpdateContext: resourceS3BucketNotificationUpdate,
		DeleteContext: resourceS3BucketNotificationDelete,

		Schema: map[string]*schema.Schema{
			"mein_feld": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceS3BucketNotificationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementieren Sie hier die Logik zum Erstellen Ihrer Ressource
	panic("not implemented create")
	return nil
}

func resourceS3BucketNotificationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementieren Sie hier die Logik zum Lesen Ihrer Ressource
	panic("not implemented read")
	return nil
}

func resourceS3BucketNotificationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementieren Sie hier die Logik zum Aktualisieren Ihrer Ressource
	panic("not implemented update")
	return nil
}

func resourceS3BucketNotificationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementieren Sie hier die Logik zum Löschen Ihrer Ressource
	panic("not implemented delete")
	return nil
}
