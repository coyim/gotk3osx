package cairo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.cairo_gobject_subpixel_order_get_type()), marshalSubpixelOrder},
		{glib.Type(C.cairo_gobject_hint_style_get_type()), marshalHintStyle},
		{glib.Type(C.cairo_gobject_hint_metrics_get_type()), marshalHintMetrics},

		// Boxed
		{glib.Type(C.cairo_gobject_font_options_get_type()), marshalFontOptions},
	}
	glib.RegisterGValueMarshalers(tm)
}

// SubpixelOrder is a representation of Cairo's cairo_subpixel_order_t.
type SubpixelOrder int

const (
	SUBPIXEL_ORDER_DEFAULT SubpixelOrder = C.CAIRO_SUBPIXEL_ORDER_DEFAULT
	SUBPIXEL_ORDER_RGB     SubpixelOrder = C.CAIRO_SUBPIXEL_ORDER_RGB
	SUBPIXEL_ORDER_BGR     SubpixelOrder = C.CAIRO_SUBPIXEL_ORDER_BGR
	SUBPIXEL_ORDER_VRGB    SubpixelOrder = C.CAIRO_SUBPIXEL_ORDER_VRGB
	SUBPIXEL_ORDER_VBGR    SubpixelOrder = C.CAIRO_SUBPIXEL_ORDER_VBGR
)

func marshalSubpixelOrder(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SubpixelOrder(c), nil
}

// HintStyle is a representation of Cairo's cairo_hint_style_t.
type HintStyle int

const (
	HINT_STYLE_DEFAULT HintStyle = C.CAIRO_HINT_STYLE_DEFAULT
	HINT_STYLE_NONE    HintStyle = C.CAIRO_HINT_STYLE_NONE
	HINT_STYLE_SLIGHT  HintStyle = C.CAIRO_HINT_STYLE_SLIGHT
	HINT_STYLE_MEDIUM  HintStyle = C.CAIRO_HINT_STYLE_MEDIUM
	HINT_STYLE_FULL    HintStyle = C.CAIRO_HINT_STYLE_FULL
)

func marshalHintStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return HintStyle(c), nil
}

// HintMetrics is a representation of Cairo's cairo_hint_metrics_t.
type HintMetrics int

const (
	HINT_METRICS_DEFAULT HintMetrics = C.CAIRO_HINT_METRICS_DEFAULT
	HINT_METRICS_OFF     HintMetrics = C.CAIRO_HINT_METRICS_OFF
	HINT_METRICS_ON      HintMetrics = C.CAIRO_HINT_METRICS_ON
)

func marshalHintMetrics(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return HintMetrics(c), nil
}

// FontOptions is a representation of Cairo's cairo_font_options_t.
type FontOptions struct {
	native *C.cairo_font_options_t
}

func marshalFontOptions(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &FontOptions{
		native: (*C.cairo_font_options_t)(unsafe.Pointer(c)),
	}, nil
}

// CreatFontOptions is a wrapper around cairo_font_options_create().
func CreateFontOptions() *FontOptions {
	native := C.cairo_font_options_create()

	opts := &FontOptions{native}
	runtime.SetFinalizer(opts, func(v *FontOptions) { glib.FinalizerStrategy(v.destroy) })

	return opts
}

func (o *FontOptions) destroy() {
	C.cairo_font_options_destroy(o.native)
}

// Copy is a wrapper around cairo_font_options_copy().
func (o *FontOptions) Copy() *FontOptions {
	native := C.cairo_font_options_copy(o.native)

	opts := &FontOptions{native}
	runtime.SetFinalizer(opts, func(v *FontOptions) { glib.FinalizerStrategy(v.destroy) })

	return opts
}

// Status is a wrapper around cairo_font_options_status().
func (o *FontOptions) Status() Status {
	return Status(C.cairo_font_options_status(o.native))
}

// Merge is a wrapper around cairo_font_options_merge().
func (o *FontOptions) Merge(other *FontOptions) {
	C.cairo_font_options_merge(o.native, other.native)
}

// Hash is a wrapper around cairo_font_options_hash().
func (o *FontOptions) Hash() uint32 {
	return uint32(C.cairo_font_options_hash(o.native))
}

// Equal is a wrapper around cairo_font_options_equal().
func (o *FontOptions) Equal(other *FontOptions) bool {
	return gobool(C.cairo_font_options_equal(o.native, other.native))
}

// SetAntialias is a wrapper around cairo_font_options_set_antialias().
func (o *FontOptions) SetAntialias(antialias Antialias) {
	C.cairo_font_options_set_antialias(o.native, C.cairo_antialias_t(antialias))
}

// GetAntialias is a wrapper around cairo_font_options_get_antialias().
func (o *FontOptions) GetAntialias() Antialias {
	return Antialias(C.cairo_font_options_get_antialias(o.native))
}

// SetSubpixelOrder is a wrapper around cairo_font_options_set_subpixel_order().
func (o *FontOptions) SetSubpixelOrder(subpixelOrder SubpixelOrder) {
	C.cairo_font_options_set_subpixel_order(o.native, C.cairo_subpixel_order_t(subpixelOrder))
}

// GetSubpixelOrder is a wrapper around cairo_font_options_get_subpixel_order().
func (o *FontOptions) GetSubpixelOrder() SubpixelOrder {
	return SubpixelOrder(C.cairo_font_options_get_subpixel_order(o.native))
}

// SetHintStyle is a wrapper around cairo_font_options_set_hint_style().
func (o *FontOptions) SetHintStyle(hintStyle HintStyle) {
	C.cairo_font_options_set_hint_style(o.native, C.cairo_hint_style_t(hintStyle))
}

// GetHintStyle is a wrapper around cairo_font_options_get_hint_style().
func (o *FontOptions) GetHintStyle() HintStyle {
	return HintStyle(C.cairo_font_options_get_hint_style(o.native))
}

// SetHintMetrics is a wrapper around cairo_font_options_set_hint_metrics().
func (o *FontOptions) SetHintMetrics(hintMetrics HintMetrics) {
	C.cairo_font_options_set_hint_metrics(o.native, C.cairo_hint_metrics_t(hintMetrics))
}

// GetHintMetrics is a wrapper around cairo_font_options_get_hint_metrics().
func (o *FontOptions) GetHintMetrics() HintMetrics {
	return HintMetrics(C.cairo_font_options_get_hint_metrics(o.native))
}
