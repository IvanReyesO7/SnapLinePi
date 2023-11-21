package services

import "github.com/line/line-bot-sdk-go/v7/linebot"

func BuildFlexSnapshotMessage(imageUrl string, text string) *linebot.FlexMessage {
	flexContainer := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:       linebot.FlexComponentTypeImage,
			URL:        imageUrl,
			Size:       linebot.FlexImageSizeTypeFull,
			AspectMode: linebot.FlexImageAspectModeTypeCover,
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "Snapshot",
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXl,
				},
				&linebot.BoxComponent{
					Type:    linebot.FlexComponentTypeBox,
					Layout:  linebot.FlexBoxLayoutTypeVertical,
					Margin:  linebot.FlexComponentMarginTypeLg,
					Spacing: linebot.FlexComponentSpacingTypeSm,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Text: "Picture taken successfully",
						},
						&linebot.BoxComponent{
							Type:    linebot.FlexComponentTypeBox,
							Layout:  linebot.FlexBoxLayoutTypeBaseline,
							Spacing: linebot.FlexComponentSpacingTypeSm,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  "Time",
									Color: "#aaaaaa",
									Size:  linebot.FlexTextSizeTypeSm,
									Flex:  linebot.IntPtr(1),
								},
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  text,
									Wrap:  true,
									Color: "#666666",
									Size:  linebot.FlexTextSizeTypeSm,
									Flex:  linebot.IntPtr(5),
									Style: linebot.FlexTextStyleTypeItalic,
								},
							},
						},
					},
				},
			},
		},
	}

	flexMessage := linebot.NewFlexMessage("Snapshot Message", flexContainer)
	return flexMessage
}

func BuildFlexClipMessage(videoUrl string, previewImageUrl string) *linebot.BubbleContainer {
	hero := linebot.VideoComponent{
		Type:       linebot.FlexComponentTypeVideo,
		URL:        videoUrl,
		PreviewURL: previewImageUrl,
		AltContent: &linebot.ImageComponent{
			Type:        linebot.FlexComponentTypeImage,
			Size:        linebot.FlexImageSizeTypeFull,
			AspectRatio: "20:13",
			AspectMode:  linebot.FlexImageAspectModeTypeCover,
			URL:         previewImageUrl,
		},
		AspectRatio: "20:13",
	}

	bubble := linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Size: linebot.FlexBubbleSizeTypeMega,
		Hero: &hero,
	}

	return &bubble
}
