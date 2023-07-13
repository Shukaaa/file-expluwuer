package utils

func IsAllowedExtension(allowedExtensions []string, extension string) bool {
	if len(allowedExtensions) == 1 && allowedExtensions[0] == "*" {
		return true
	}

	for _, allowedExtension := range allowedExtensions {
		if allowedExtension == extension {
			return true
		}
	}

	return false
}

func IsBlacklistedExtension(blacklistedExtensions []string, extension string) bool {
	for _, blacklistedExtension := range blacklistedExtensions {
		if blacklistedExtension == extension {
			return true
		}
	}

	return false
}

func IsImageExtension(extension string) bool {
	imageExtensions := []string{"jpg", "jpeg", "png", "gif", "bmp", "svg", "webp", "ico", "tiff", "tif"}

	for _, imageExtension := range imageExtensions {
		if imageExtension == extension {
			return true
		}
	}

	return false
}

func IsVideoExtension(extension string) bool {
	videoExtensions := []string{"mp4", "avi", "flv", "wmv", "mov", "mkv", "webm", "m4v", "mpg", "mpeg", "3gp", "3g2", "m2v", "m4v", "svi", "mxf", "roq", "nsv", "f4v", "f4p", "f4a", "f4b"}

	for _, videoExtension := range videoExtensions {
		if videoExtension == extension {
			return true
		}
	}

	return false
}

func IsAudioExtension(extension string) bool {
	audioExtensions := []string{"mp3", "wav", "flac", "ogg", "m4a", "wma", "aac", "aiff", "alac", "pcm", "dsd", "dsf", "dff", "mka", "m3u", "pls", "opus", "weba", "caf", "amr", "mid", "midi", "m4p", "m4r", "3ga", "aif", "aifc", "cda"}

	for _, audioExtension := range audioExtensions {
		if audioExtension == extension {
			return true
		}
	}

	return false
}

func IsArchiveExtension(extension string) bool {
	archiveExtensions := []string{"zip", "rar", "7z", "tar", "gz", "gzip", "xz", "bz2", "bzip2", "z", "taz", "tbz", "tbz2", "lz", "tlz", "txz", "lzma", "tlzma", "lzma86", "lz", "lzo", "tar.lzo", "tar.lzop", "tar.gz", "tar.bz2", "tar.xz", "tar.lzma", "tlz", "tar.Z", "tar.z", "tar.zst", "tzst", "tar.Z", "tar.z", "tar.zst", "tzst", "cpio", "cpio.gz", "cpio.bz2", "cpio.xz", "cpio.lzma", "cpio.lz", "cpio.lzo", "cpio.Z", "cpio.z", "cpio.zst", "cpio.tzst", "iso", "wim", "swm", "udf", "chm", "chm"}

	for _, archiveExtension := range archiveExtensions {
		if archiveExtension == extension {
			return true
		}
	}

	return false
}

func IsExecutableExtension(extension string) bool {
	executableExtensions := []string{"exe", "msi", "bin", "sh", "bat", "cmd", "app", "apk", "ipa", "xap", "xexe", "xmsi", "xapk", "xapp", "xip", "xap", "xexe", "xmsi", "xapk", "xapp", "xip"}

	for _, executableExtension := range executableExtensions {
		if executableExtension == extension {
			return true
		}
	}

	return false
}

func IsFontExtension(extension string) bool {
	fontExtensions := []string{"ttf", "otf", "woff", "woff2", "eot", "suit", "ttc", "font", "fonts", "sui"}

	for _, fontExtension := range fontExtensions {
		if fontExtension == extension {
			return true
		}
	}

	return false
}
