package builder

import (
	"github.com/hashicorp/waypoint/builtin/docker"
)

// NixpacksImageMapper maps a nixpacks.Image to a docker.Image structure.
func NixpacksImageMapper(src *Image) *docker.Image {
	return &docker.Image{
		Image: src.Image,
		Tag:   src.Tag,

		Location: &docker.Image_Docker{},
	}
}
