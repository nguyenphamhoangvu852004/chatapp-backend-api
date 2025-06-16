package middleware

import (
	"chapapp-backend-api/global"
	"context"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadProfileAccountToCloudinary() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy file từ form-data
		cloud := global.Config.Cloudinary
		avatar, err := c.FormFile("avatar")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}
		cover, err := c.FormFile("cover")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}
		// đọc file hinh anh avatar
		avatarFile, err := avatar.Open()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
			return
		}
		defer avatarFile.Close()

		// đọc file hinh anh đại diện
		coverFile, err := cover.Open()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
			return
		}
		defer coverFile.Close()

		// Khởi tạo Cloudinary
		cld, err := cloudinary.NewFromParams(cloud.CloudName, cloud.ApiKey, cloud.ApiSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary init failed"})
			return
		}

		// Upload ảnh đại diện
		uploadAvatarRs, err := cld.Upload.Upload(context.Background(), avatarFile, uploader.UploadParams{
			Folder:   "chatapp",
			PublicID: uuid.New().String(),
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Upload failed", "detail": err.Error()})
			return
		}

		// Upload ảnh bìa
		uploadCoverRs, err := cld.Upload.Upload(context.Background(), coverFile, uploader.UploadParams{
			Folder:   "chatapp",
			PublicID: uuid.New().String(),
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Upload failed", "detail": err.Error()})
			return
		}

		// Gắn URL file vào context để controller xài
		c.Set("avatarUrl", uploadAvatarRs.SecureURL)
		c.Set("coverUrl", uploadCoverRs.SecureURL)

		// Tiếp tục middleware chain
		c.Next()
	}
}

func UploadGroupAvatarToCloundinary() gin.HandlerFunc {
	return func (c *gin.Context)  {
	cloud := global.Config.Cloudinary

	avatar, err := c.FormFile("avatar")	

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// đọc file hinh anh avatar
	avatarFile, err := avatar.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
		return
	}
	defer avatarFile.Close()

	// Khởi tạo Cloudinary
	cld, err := cloudinary.NewFromParams(cloud.CloudName, cloud.ApiKey, cloud.ApiSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary init failed"})
		return
	}

	// Upload ảnh đại diện	
	uploadAvatarRs, err := cld.Upload.Upload(context.Background(), avatarFile, uploader.UploadParams{
		Folder:   "chatapp",
		PublicID: uuid.New().String(),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Upload failed", "detail": err.Error()})
		return
	}

	// Gắn URL file vào context để controller xài
	c.Set("avatarUrl", uploadAvatarRs.SecureURL)
	c.Next()
	}
}
func ModifyUploadGroupAvatarToCloundinary() gin.HandlerFunc {
	return func (c *gin.Context)  {
	cloud := global.Config.Cloudinary

	avatar, err := c.FormFile("avatar")	

	if err != nil {
		c.Next()
		return
	}

	// đọc file hinh anh avatar
	avatarFile, err := avatar.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
		return
	}
	defer avatarFile.Close()

	// Khởi tạo Cloudinary
	cld, err := cloudinary.NewFromParams(cloud.CloudName, cloud.ApiKey, cloud.ApiSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary init failed"})
		return
	}

	// Upload ảnh đại diện	
	uploadAvatarRs, err := cld.Upload.Upload(context.Background(), avatarFile, uploader.UploadParams{
		Folder:   "chatapp",
		PublicID: uuid.New().String(),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Upload failed", "detail": err.Error()})
		return
	}

	// Gắn URL file vào context để controller xài
	c.Set("avatarUrl", uploadAvatarRs.SecureURL)
	c.Next()
	}
}

