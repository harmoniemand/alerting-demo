package notifications

type Notification struct {
	/* The sender of the notification! */
	Sender string `json:"sender"`

	/* The type of the notification!
	 * This can be used to filter notifications by type!
	 * Type should be debug, info, warning, error!
	 */
	Type string `json:"type" validate:"required"`

	/* The name of the notification!
	 * This can be used to filter notifications by name!
	 * The name should be a very short description of the notification!
	 */
	Name string `json:"name" validate:"required"`

	/* The description of the notification!
	 * This can be used to easyly understand the notification!
	 * Should not be too long but very descriptive!
	 */
	Description string `json:"description" validate:"required"`
}
