package middleware

import (
	"chapapp-backend-api/global"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadMediaToCloudinary() gin.HandlerFunc {
	return func(c *gin.Context) {
		cloud := global.Config.Cloudinary

		// Lấy tất cả file images và videos từ form-data
		form, err := c.MultipartForm()
		if err != nil {
			c.Next()
			return
		}

		imageFiles := form.File["images"]
		videoFiles := form.File["videos"]

		cld, err := cloudinary.NewFromParams(cloud.CloudName, cloud.ApiKey, cloud.ApiSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary init failed"})
			return
		}

		imageUrls := []string{}
		for _, file := range imageFiles {
			url, err := uploadToCloudinary(cld, file, "image")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Image upload failed", "detail": err.Error()})
				return
			}
			imageUrls = append(imageUrls, url)
		}

		videoUrls := []string{}
		for _, file := range videoFiles {
			url, err := uploadToCloudinary(cld, file, "video")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Video upload failed", "detail": err.Error()})
				return
			}
			videoUrls = append(videoUrls, url)
		}

		// Lưu URLs vào context
		c.Set("imageUrls", imageUrls)
		c.Set("videoUrls", videoUrls)

		c.Next()
	}
}
func uploadToCloudinary(cld *cloudinary.Cloudinary, file *multipart.FileHeader, resourceType string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("cannot open file: %w", err)
	}
	defer src.Close()

	uploadParams := uploader.UploadParams{
		Folder:       "chatapp",
		PublicID:     uuid.New().String(),
		ResourceType: resourceType,
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), src, uploadParams)
	if err != nil {
		return "", fmt.Errorf("upload error: %w", err)
	}

	return uploadResult.SecureURL, nil
}
func UploadProfileAccountToCloudinary() gin.HandlerFunc {
	return func(c *gin.Context) {
		cloud := global.Config.Cloudinary

		cld, err := cloudinary.NewFromParams(cloud.CloudName, cloud.ApiKey, cloud.ApiSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary init failed"})
			return
		}

		// Upload avatar nếu có
		if avatar, err := c.FormFile("avatar"); err == nil {
			avatarFile, err := avatar.Open()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot open avatar file"})
				return
			}
			defer avatarFile.Close()

			uploadAvatarRs, err := cld.Upload.Upload(context.Background(), avatarFile, uploader.UploadParams{
				Folder:   "chatapp",
				PublicID: uuid.New().String(),
			})
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Upload avatar failed", "detail": err.Error()})
				return
			}
			c.Set("avatarUrl", uploadAvatarRs.SecureURL)
		}

		// Upload cover nếu có
		if cover, err := c.FormFile("cover"); err == nil {
			coverFile, err := cover.Open()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot open cover file"})
				return
			}
			defer coverFile.Close()

			uploadCoverRs, err := cld.Upload.Upload(context.Background(), coverFile, uploader.UploadParams{
				Folder:   "chatapp",
				PublicID: uuid.New().String(),
			})
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Upload cover failed", "detail": err.Error()})
				return
			}
			c.Set("coverUrl", uploadCoverRs.SecureURL)
		}

		// Tiếp tục xử lý
		c.Next()
	}
}

func UploadGroupAvatarToCloundinary() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	return func(c *gin.Context) {
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
