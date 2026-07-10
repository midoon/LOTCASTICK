CREATE TABLE trade_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    trade_id UUID NOT NULL,
    user_id UUID NOT NULL,

    storage_key VARCHAR(500) NOT NULL,
    original_filename VARCHAR(255) NOT NULL,
    mime_type VARCHAR(50) NOT NULL,

    file_size_bytes INTEGER NOT NULL,

    label VARCHAR(100),

    sort_order SMALLINT NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_trade_images_trade_id
        FOREIGN KEY (trade_id)
        REFERENCES trades(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_trade_images_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_trade_images_storage_key
        UNIQUE (storage_key),

    CONSTRAINT chk_trade_images_file_size
        CHECK (
            file_size_bytes > 0
            AND file_size_bytes <= 5242880
        ),

    CONSTRAINT chk_trade_images_mime_type
        CHECK (
            mime_type IN (
                'image/jpeg',
                'image/png',
                'image/webp'
            )
        )
);

CREATE INDEX idx_trade_images_trade_id
    ON trade_images (trade_id)
    WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX idx_trade_images_storage_key
    ON trade_images (storage_key);